#!/bin/bash
GO=schemas.go
SCMDIR="doc/EDDN/schemas/"
pushd $SCMDIR > /dev/null
SCMSRCS=$(ls -1 *-v*.json)
popd > /dev/null
for scm in $SCMSRCS; do
	nm=$(echo $scm | sed -E 's/(.*)-.*/\1/')
	scnm=$scnm" "$nm
done

echo "// generated with genschemas.sh" > $GO
echo "package eddn" >> $GO
echo "const (" >> $GO
sep=" ScmId = iota"
for scm in $scnm; do
	echo "	S"$scm$sep >> $GO
	sep=""
done
echo ")" >> $GO
echo "var ScmURLs = []string{" >> $GO
for fs in $SCMSRCS; do
	echo -n "	\"" >> $GO
	egrep '^[ \t]*"id"[ \t]*:' $SCMDIR$fs \
	| sed -E 's/.*(https[^#]*).*/\1/' \
	| tr -d \\n >> $GO
	echo "\"," >> $GO
done
echo "}" >> $GO
echo
echo "var ScmDefs = []string{" >> $GO
for fs in $SCMSRCS; do
	echo -n "\`" >> $GO
	cat $SCMDIR$fs >> $GO
	echo "\`," >> $GO
done
echo "}" >> $GO

gofmt $GO > tmp && mv tmp $GO
