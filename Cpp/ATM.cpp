#pragma once

#include <string>

struct card_inserted {
    std::string account;
};

class ATM
{
    messaging::receiver incoming;
    messaging::sender bank;
    messaging::sender interface_hardware;
    
    void (ATM::*state)();

    std::string account;
    std::string pin;

    void waiting_for_card() {
        interface_hardware.send(display_enter_card());
        incoming.wait()
            .handle<card_inserted>(
                [&](card_inserted const& msg) {
                    account = msg.account;
                    pin = "";
                    interface_hardware.send(display_enter_pin());
                    state = &ATM::getting_pin;
                }
            );
    }

    void getting_pin();

public:
    void run() {
        state = &ATM::waiting_for_card;
        try {
            for (;;) {
                (this->*state)();
            }
        } catch (messaging::close_queue const&) {

        }
    }
};

void ATM::getting_pin() {
    incoming.wait()
        .handle<digit_pressed>(
            [&](digit_pressed const& msg) {
                unsigned const pin_length = 4;
                pin += msg.digit;
                if (pin.length() == pin_length) {
                    bank.send(verify_pin(account, pin, incoming));
                    state = &ATM::verifying_pin;
                }
            }
        )
        .handle<clear_last_pressed>(
            [&](clear_last_pressed const& msg) {
                if (!pin_empty()) {
                    pin.resize(pin.length() - 1);
                }
            }                            
        )
        .handle<cancel_pressed>(
            [&](cancel_pressed const& msg) {
                state = &ATM::done_processing;
            }                        
        );
}
