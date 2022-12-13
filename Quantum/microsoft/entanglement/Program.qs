namespace entanglement {
    open Microsoft.Quantum.Canon;
    open Microsoft.Quantum.Intrinsic;
    
    operation SetQubitState(desired : Result, q1 : Qubit) : Unit {
        if (desired != M(q1)) {
            X(q1);
        }
    }

    @EntryPoint()
    operation TestBellState(count : Int, initial : Result) : (Int, Int, Int) {
        mutable numOnes = 0;
        mutable agree = 0;

        using ((qubit0, qubit1) = (Qubit(), Qubit())) {

            for (test in 1..count) {
                SetQubitState(initial, qubit0);
                SetQubitState(Zero, qubit1);
                
                H(qubit0);
                CNOT(qubit0, qubit1);
                let res = M(qubit0);

                if (M(qubit1) == res) {
                    set agree += 1;
                }

                if (res == One) {
                    set numOnes += 1;
                }
            }

            SetQubitState(Zero, qubit0);
            SetQubitState(Zero, qubit1);
        }
    
        Message("Test results (# of 0s, # of 1s): ");
        return (count - numOnes, numOnes, agree);
    }
}
