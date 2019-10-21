
go build -o ./clok -ldflags "-w" ./clock
go build -o ./clockwall -ldflags "-w" ./main.go

ZONEINFO="C:\\Users\\da00440056\\Documents\\cygwin\\usr\\share\\zoneinfo"

TZ=Asia/Tokyo ./clok -p 8001 -z $ZONEINFO & \
TZ=US/Alaska ./clok -p 8002 -z $ZONEINFO & \
TZ=US/Eastern ./clok -p 8003 -z $ZONEINFO & \
TZ=US/Pacific ./clok -p 8004 -z $ZONEINFO & \
TZ=Europe/Berlin ./clok -p 8005 -z $ZONEINFO & \
TZ=Europe/London ./clok -p 8006 -z $ZONEINFO & \
TZ=Europe/Madrid ./clok -p 8007 -z $ZONEINFO & \
TZ=Europe/Moscow ./clok -p 8008 -z $ZONEINFO & \
TZ=Asia/Dubai ./clok -p 8009 -z $ZONEINFO & \
TZ=Hongkong ./clok -p 8010 -z $ZONEINFO & \
TZ=Asia/Tel_Aviv ./clok -p 8011 -z $ZONEINFO & \
TZ=ROC ./clok -p 8012 -z $ZONEINFO & \
TZ=Australia/Sydney ./clok -p 8013 -z $ZONEINFO &

./clockwall \
    "Japan, Tokyo=localhost:8001" \
    "Alaska, U.S.A.=localhost:8002" \
    "New York, U.S.A.=localhost:8003" \
    "California, U.S.A.=localhost:8004" \
    "Berlin, Germany=localhost:8005" \
    "London, U.K.=localhost:8006" \
    "Madrid, Spain=localhost:8007" \
    "Moscow, Russia=localhost:8008" \
    "Dubai, U.A.E.=localhost:8009" \
    "Hong Kong=localhost:8010" \
    "Tel Aviv, Israel=localhost:8011" \
    "ROC=localhost:8012" \
    "Sydney, Australia=localhost:8013"
