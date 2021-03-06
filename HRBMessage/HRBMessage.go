package HRBMessage

type TcpHeader int

const (
	MSG  TcpHeader = 0
	ECHO TcpHeader = 1
	ACC  TcpHeader = 2
	REQ  TcpHeader = 3
	FWD  TcpHeader = 4
	BIN  TcpHeader = 5
	Stat  TcpHeader = 6
)

type Message interface {
	GetHeaderType() TcpHeader
	GetHashData() string
	GetData() string
	GetId() string
	GetRound() int
	GetSenderId() string
	SetSenderIdNull()
	SetSenderId(string)
}

type FWDStruct struct {
	Header TcpHeader
	Data   string
	HashData string
	Round int
	Id string
	SenderId string
}

func (d FWDStruct) SetDataNull(){

}

type MSGStruct struct {
	Header TcpHeader
	Data   string
	HashData string
	Round int
	Id string
	SenderId string
}

func (d MSGStruct) SetDataNull(){
	d.Data = ""
}

type ECHOStruct struct {
	Header TcpHeader
	HashData string
	Data string
	Round int
	Id string
	SenderId string
}

func (d ECHOStruct) SetDataNull(){
	(&d).Data = ""
}

type ACCStruct struct {
	Header TcpHeader
	HashData string
	Data string
	Round int
	Id string
	SenderId string
}

func (d ACCStruct) SetDataNull(){

}

type REQStruct struct {
	Header TcpHeader
	HashData string
	Round int
	Id string
	SenderId string
}

func (d REQStruct) SetDataNull(){

}

type PrepareSend struct {
	M      Message
	SendTo string
	Stat   Stats
}

type Binary struct {
	Header TcpHeader
	HashData string
	Round int
	SenderId string
	Id string
}

func (d Binary) SetDataNull(){

}

type StatStruct struct {
	Header TcpHeader
	Round int
	Id string
}


func (b StatStruct) GetRecSend() {
}

func (b StatStruct) GetHeaderType() TcpHeader {
	return b.Header
}

func (b StatStruct) GetHashData() string {
	return ""
}

func (b StatStruct) GetData() string {
	return ""
}

func (b StatStruct) GetId() string{
	return b.Id
}

func (b StatStruct) GetRound() int {
	return b.Round
}

func (b StatStruct) GetSenderId() string{
	return ""
}

func (b StatStruct)SetSenderIdNull() {

}

func (b StatStruct) SetSenderId(string) {

}

func (b StatStruct) SetDataNull() {

}


func (b Binary) GetHeaderType() TcpHeader {
	return b.Header
}

func (b Binary) GetHashData() string {
	return b.HashData
}

func (b Binary) GetData() string {
	return ""
}

func (b Binary) GetId() string{
	return b.Id
}

func (b Binary) GetRound() int {
	return b.Round
}

func (b Binary) GetSenderId() string {
	return b.SenderId
}

func (b Binary)SetSenderIdNull() {

}

func (b Binary) SetSenderId(string) {

}


/*
Implement the interface
 */

func (m FWDStruct) GetHeaderType() TcpHeader{
	return m.Header
}

func (m MSGStruct) GetHeaderType() TcpHeader{
	return m.Header
}

func (m ECHOStruct) GetHeaderType() TcpHeader{
	return m.Header
}

func (m ACCStruct) GetHeaderType() TcpHeader{
	return m.Header
}

func (m REQStruct) GetHeaderType() TcpHeader{
	return m.Header
}



/*
Return the Value contained in the message for real message
 */

func (m FWDStruct) GetData() string{
	return m.Data
}

func (m MSGStruct) GetData() string {
	return m.Data
}

func (m ECHOStruct) GetData() string {
	return m.Data
}

func (m ACCStruct) GetData() string {
	return m.Data
}

func (m REQStruct) GetData() string {
	return ""
}


/*
Return the Value contained in the message for real message
*/

func (m FWDStruct) GetHashData() string{
	return m.HashData
}

func (m MSGStruct) GetHashData() string{
	return m.HashData
}

func (m ECHOStruct) GetHashData() string{
	return m.HashData
}

func (m ACCStruct) GetHashData() string{
	return m.HashData
}

func (m REQStruct) GetHashData() string{
	return m.HashData
}



/*
Get the senderID of this message
*/

func (m FWDStruct) GetId() string{
	return m.Id
}

func (m MSGStruct) GetId() string{
	return m.Id
}

func (m ECHOStruct) GetId() string{
	return m.Id
}

func (m ACCStruct) GetId() string{
	return m.Id
}

func (m REQStruct) GetId() string{
	return m.Id
}


/*
Get the Round
*/

func (m FWDStruct) GetRound() int{
	return m.Round
}

func (m MSGStruct) GetRound() int{
	return m.Round
}

func (m ECHOStruct) GetRound() int{
	return m.Round
}

func (m ACCStruct) GetRound() int{
	return m.Round
}

func (m REQStruct) GetRound() int{
	return m.Round
}



func (m FWDStruct) GetSenderId() string{
	return m.SenderId
}

func (m MSGStruct) GetSenderId() string{
	return m.SenderId
}

func (m ECHOStruct) GetSenderId() string{
	return m.SenderId
}

func (m ACCStruct) GetSenderId() string{
	return m.SenderId
}

func (m REQStruct) GetSenderId() string{
	return m.SenderId
}


func (m FWDStruct) SetSenderIdNull() {
	m.SenderId = ""
}

func (m MSGStruct) SetSenderIdNull() {
	m.SenderId = ""
}

func (m ECHOStruct) SetSenderIdNull() {
	m.SenderId = ""
}

func (m ACCStruct) SetSenderIdNull() {
	m.SenderId = ""
}

func (m REQStruct) SetSenderIdNull() {
	m.SenderId = ""
}




func (m FWDStruct) SetSenderId(id string) {
	m.SenderId = id
}

func (m MSGStruct) SetSenderId(id string) {
	m.SenderId = id
}

func (m ECHOStruct) SetSenderId(id string) {
	m.SenderId = id
}

func (m ACCStruct) SetSenderId(id string) {
	m.SenderId = id
}

func (m REQStruct) SetSenderId(id string) {
	m.SenderId = id
}
