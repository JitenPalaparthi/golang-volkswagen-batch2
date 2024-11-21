- When you run an application ?

- A process is created and the binary code is loaded in that process
- 

java source code -> javac -> java bytecode -> jvm (JIT)
clojure           > clojurec /
scala             > scala   /
kotlin            > kotlin /


c,c++ --> compiler(gcc/llvm) -> FE(llvm/gcc) -> BE(llvm/gcc) --> assembler - linker - binary

go    --> go compiler (SSA)  -> assembler(plan9)(.o) - linker -> binary

SSA-> Static Single Assignment form -> optimization (IR)