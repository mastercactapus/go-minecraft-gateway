package main

type KeepAlive struct {
	//id 0x00

	KeepAliveID int32
}

type JoinGame struct {
	//id 0x01

	EntityID int32
	Gamemode int8
	Dimension uint8
	Difficulty uint8
	MaxPlayers uint8
	LevelType string
}

type ChatMessage struct {
	//id 0x02

	JSONData string
}

type TimeUpdate struct {
	//id 0x03

	AgeOfWorld int64
	TimeOfDay int64
}

type EntityEquipment struct {
	//id 0x04

	EntityID int32
	Slot int16
	// Item Slot
}

type SpawnPosition struct {
	//id 0x05

	X int32
	Y int32
	Z int32
}

type UpdateHealth struct {
	//id 0x06

	Health float32
	Food int16
	FoodSaturation float32
}

type Respawn struct {
	//id 0x07

	Dimension int32
	Difficulty uint8
	Gamemode uint8
	LevelType string
}

type PlayerPositionAndLook struct {
	//id 0x08

	X float64
	Y float64
	Z float64
	Yaw float32
	Pitch float32
	OnGround bool
}

type HeldItemChange struct {
	//id 0x09

	Slot int8
}

type UseBed struct {
	//id 0x0a

	EntityID int32
	X int32
	Y uint8
	Z int32
}


type Animation struct {
	//id 0x0b

	EntityID int32
	Animation uint8 //Anim<type> const
}



type SpawnPlayer struct {
	//id 0x0b

	EntityID int32
	PlayerUUID string
	PlayerName string
	DataCount int32
	Data []PropertyData
	X int32
	Y int32
	Z int32
	Yaw int8
	Pitch int8
	CurrentItem int16
	Metadata Metadata
}
