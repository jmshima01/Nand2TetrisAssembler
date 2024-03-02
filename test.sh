#!bin/sh
make
for i in test/*.asm; do
    ./Assembler $i
done


 