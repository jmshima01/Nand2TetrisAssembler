#!bin/sh
make
chmod +x Assembler
for i in asm/*.asm; 
do
    ./Assembler $i
done

for i in tests/*; do
    X=$(echo "$i" | sed "s/tests\//asm\//")
    D=$(diff --strip-trailing-cr $i $X)
    if [$D == ""] 
    then
        echo "Passed $i"
    fi
done