// Quantum Trial

//namespace quantum {
//    open Microsoft.Quantum.Canon;
//    open Microsoft.Quantum.Intrinsic;
//    
//
//    @EntryPoint()
//    operation SayHello() : Unit {
//        Message("Hello quantum world!");
//    }
//}

namespace Qrng {
    open Microsoft.Quantum.Convert;
    open Microsoft.Quantum.Math;
    open Microsoft.Quantum.Measurement;
    open Microsoft.Quantum.Canon;
    open Microsoft.Quantum.Intrinsic;

    operation SampleQuantumRandomNumberGenerator() : Result {
        using (q = Qubit()) {
            H(q);
            return MResetZ(q);
        }
    }

    operation SampleRandomNumberInRange(max : Int) : Int {
        mutable bits = new Result[0];
        for (idxBit in 1..BitSizeI(max)) {
            set bits += [SampleQuantumRandomNumberGenerator()];
        }
        let sample = ResultArrayAsInt(bits);
        return sample > max
               ? SampleRandomNumberInRange(max)
               | sample;
    }

    @EntryPoint()
    operation SmapleRandomNumber() : Int {
        let max = 50;
        Message($"Sampling a random number between 0 and {max}: ");
        return SampleRandomNumberInRange(max);
    }
}
