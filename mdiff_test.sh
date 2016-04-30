#!/bin/bash

echo "Test case 1."
rm _tmp.txt
rm _tmp_.txt

for i in {1..10000}
do
	echo $i >> _tmp.txt
done
cp _tmp.txt _tmp_.txt
sed -i -e 's/\(111\)/_\1/g' _tmp_.txt

go run mdiff.go _tmp.txt _tmp_.txt

echo "Test case 2."
rm Diff.html
rm Diff_.html

curl -s -o Diff.html https://ja.wikipedia.org/wiki/Diff
cp Diff.html Diff_.html

sed -i -e '100d' Diff_.html
sed -i -e '200d' Diff_.html
sed -i -e '300d' Diff_.html

sed -i -e '150i\
	AAA
' Diff_.html
sed -i -e '250i\
	BBB
' Diff_.html
sed -i -e '350i\
	CCC
' Diff_.html

sed -i -e 's/https/spdy/g' Diff_.html

go run mdiff.go Diff.html Diff_.html

echo "Test case 3."
for i in {0..10000}
do
	echo $RANDOM >> rand.txt
done

sort rand.txt > rand_.txt

go run mdiff.go rand.txt rand_.txt

