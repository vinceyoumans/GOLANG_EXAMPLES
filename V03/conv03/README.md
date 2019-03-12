//  Will build structs with data
// then save to bolt Storm DB
// Then save to Spread sheet

go get github.com/360EntSecGroup-Skylar/excelize

go get -u github.com/asdine/storm

go run main.go lib01.go lib02.go lib03.go lib10.go struct01.go
