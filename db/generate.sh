root="db/"
in="template.map"

plural="blueprints"
singular="Blueprint"
genny -in=$root$in -out=$root$plural-funcs.go -pkg=db gen "groupName=$plural itemName=$singular fileName=\"$plural\""
echo $root$plural-funcs.go created

plural="categoryIDs"
singular="CategoryID"
genny -in=$root$in -out=$root$plural-funcs.go -pkg=db gen "groupName=$plural itemName=$singular fileName=\"$plural\""
echo $root$plural-funcs.go created

plural="certificates"
singular="Certificate"
genny -in=$root$in -out=$root$plural-funcs.go -pkg=db gen "groupName=$plural itemName=$singular fileName=\"$plural\""
echo $root$plural-funcs.go created

plural="graphicIDs"
singular="GraphicID"
genny -in=$root$in -out=$root$plural-funcs.go -pkg=db gen "groupName=$plural itemName=$singular fileName=\"$plural\""
echo $root$plural-funcs.go created

plural="groupIDs"
singular="GroupID"
genny -in=$root$in -out=$root$plural-funcs.go -pkg=db gen "groupName=$plural itemName=$singular fileName=\"$plural\""
echo $root$plural-funcs.go created

plural="iconIDs"
singular="IconID"
genny -in=$root$in -out=$root$plural-funcs.go -pkg=db gen "groupName=$plural itemName=$singular fileName=\"$plural\""
echo $root$plural-funcs.go created

plural="skins"
singular="Skin"
genny -in=$root$in -out=$root$plural-funcs.go -pkg=db gen "groupName=$plural itemName=$singular fileName=\"$plural\""
echo $root$plural-funcs.go created

plural="typeIDs"
singular="TypeID"
genny -in=$root$in -out=$root$plural-funcs.go -pkg=db gen "groupName=$plural itemName=$singular fileName=\"$plural\""
echo $root$plural-funcs.go created

plural="marketGroups"
singular="MarketGroup"
fileName="invMarketGroups"
primaryKey="marketGroupID"
genny -in=$root$in -out=$root$plural-funcs.go -pkg=db gen "groupName=$plural itemName=$singular fileName=\"$fileName\"" primaryKey=$primaryKey
echo $root$plural-funcs.go created

go install