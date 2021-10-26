#!/bin/bash
GO=schemas.go
SCMDIR="doc/EDDN/schemas/"
pushd $SCMDIR > /dev/null
SCMSRCS=$(ls -1 *-v*.json)
popd > /dev/null
scmno=0
for scm in $SCMSRCS; do
	nm=$(echo $scm | sed -E 's/(.*)-.*/\1/')
	scnm=$scnm" "$nm
	scmno=$(($scmno + 1))
done

echo "// generated with genschemas.sh" > $GO
echo -e "package eddnc\n" >> $GO
echo -e "const ScmNo = "$scmno >> $GO
echo -e "//go:generate stringer -type ScmID" >> $GO
echo "const (" >> $GO
sep=" ScmID = iota"
for scm in $scnm; do
	echo "	S"$scm$sep >> $GO
	sep=""
done
echo -e ")\n" >> $GO

for fs in $SCMSRCS; do
    scmurls=$scmurls" "$(egrep '^[ \t]*"id"[ \t]*:' $SCMDIR$fs \
	                         | sed -E 's/.*(https[^#]*).*/\1/')
done

echo "var ScmURLs = []string{" >> $GO
for u in $scmurls; do
	echo "	\""$u"\"," >> $GO
done
echo -e "}\n" >> $GO
idx=0
echo "var ScmMap = map[string]ScmID{" >> $GO
for u in $scmurls; do
    echo -e "\t\""$u"\": "$idx"," >> $GO
    idx=$(($idx + 1))
done
echo -e "}\n" >> $GO
echo "var ScmDefs = []string{" >> $GO
for fs in $SCMSRCS; do
	echo -n "\`" >> $GO
	cat $SCMDIR$fs >> $GO
	echo "\`," >> $GO
done
echo "}" >> $GO

gofmt $GO > tmp && mv tmp $GO
