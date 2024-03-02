#!/bin/sh
cd ~/Documents/nand2tetris/nand2tetris/tools
for a in ~/Documents/nand2tetris/Nand2TetrisAssembler/test/*.asm; do
    source Assembler.sh "$a"
done
    
