# First Go Servic 🎉


## Api

| Name | Url | Description |
| ------ | ------ | ------ |
| GET | owner/info | get info about owner |
| PUT | owner/info | update info |
| GET | owner/bank | get all account |
| POST | owner/bank | create account |
| PUT | owner/bank | update account |
| DELETE | owner/bank | eelete account |

### GET owner/info
Response
```json
{
	"name" : "Sakarn Bantadjun",
	"address" : "เลขที่ 88/8888 ตำบล/แขวน ลาดกะบัง อำเภอ/เขต ลาดกระบัง จังหวัด กรุงเทพ",
	"tal" : "0888888888",
	"waterBill" : 20,
	"electricityBill" : 3,
	"servicePerDat" : 300
}
```

### PUT owner/info
Request
```json
{
	"name" : "Sakarn Bantadjun",
	"address" : "เลขที่ 88/8888 ตำบล/แขวน ลาดกะบัง อำเภอ/เขต ลาดกระบัง จังหวัด กรุงเทพ",
	"tal" : "0888888888",
	"waterBill" : 20,
	"electricityBill" : 3,
	"servicePerDat" : 300
}
```
Response
```json
true
```

### GET owner/bank
```json
[
	{
		"name" : "กรุงเทพ",
		"number" : "1234567890"
	},
	{
		"name" : "ไทยประกันสังคม",
		"number" : "1234567890"
	}
]
```

### POST owner/bank
Repuest
```json
{
	"id" : 2,
	"name" : "กรุงเทพ",
	"number" : "000000000"
}
```
Response
```json
true
```

### PUT owner/bank
Repuest
```json
{
  "_comment1":"ID is a PK to edit"
	"id" : 2,
	"name" : "กรุงเทพ",
	"number" : "000000000"
}
```
Response
```json
true
```

### Delete owner/bank
Request
```json
{
	"name" : "กรุงเทพ"
}
```
Response
```json
true
```
