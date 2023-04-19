package model

type Config struct {
	Server struct {
		Port          int    `yaml:"port"`
		Host          string `yaml:"host"`
		MaxConnection int    `yaml:"maxconnection"`
	} `yaml:"server"`

	Database struct {
		Dbname   string `yaml:"dbname"`
		User     string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

type Params struct {
	Object []Object `xml:"PARAM"`
}

type Param struct {
	ID          int64  `xml:"ID,attr"`
	OBJECTID    int64  `xml:"OBJECTID,attr"`
	CHANGEID    int64  `xml:"CHANGEID,attr"`
	CHANGEIDEND int64  `xml:"CHANGEIDEND,attr"`
	TYPEID      int    `xml:"TYPEID,attr"`
	VALUE       string `xml:"VALUE,attr"`
	UPDATEDATE  string `xml:"UPDATEDATE,attr"`
	STARTDATE   string `xml:"STARTDATE,attr"`
	ENDDATE     string `xml:"ENDDATE,attr"`
}

type Object struct {
	ID         int64  `xml:"ID,attr"`
	OBJECTID   int64  `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   int64  `xml:"CHANGEID,attr"`
	NAME       string `xml:"NAME,attr"`
	TYPENAME   string `xml:"TYPENAME,attr"`
	LEVEL      int    `xml:"LEVEL,attr"`
	OPERTYPEID int    `xml:"OPERTYPEID,attr"`
	PREVID     int    `xml:"PREVID,attr"`
	NEXTID     int    `xml:"NEXTID,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTUAL   bool   `xml:"ISACTUAL,attr"`
	ISACTIVE   bool   `xml:"ISACTIVE,attr"`
}

type AbbressObjects struct {
	Object []Object `xml:"OBJECT"`
}

var (
	DataAddresses = `
<?xml version="1.0" encoding="UTF-8"?>
<ADDRESSOBJECTS>
<OBJECT ID="1231506" OBJECTID="994428" OBJECTGUID="1ea1af38-e63f-458e-a848-63195944f8e6" CHANGEID="2725312" NAME="Наугорский" TYPENAME="пер" LEVEL="8" OPERTYPEID="1" PREVID="0" NEXTID="0" UPDATEDATE="2011-09-14" STARTDATE="1900-01-01" ENDDATE="2079-06-06" ISACTUAL="1" ISACTIVE="1"/>
<OBJECT ID="1239735" OBJECTID="1001862" OBJECTGUID="98c33627-cb6a-4fd2-9427-f932aa64cf9a" CHANGEID="2743437" NAME="Образцовская" TYPENAME="ул" LEVEL="8" OPERTYPEID="10" PREVID="0" NEXTID="0" UPDATEDATE="2017-02-08" STARTDATE="2017-02-08" ENDDATE="2079-06-06" ISACTUAL="1" ISACTIVE="1"/>
<OBJECT ID="1243644" OBJECTID="1005477" OBJECTGUID="b5064c0d-3267-4592-8ef9-ff0171ca656f" CHANGEID="2752031" NAME="Северная" TYPENAME="ул" LEVEL="8" OPERTYPEID="1" PREVID="0" NEXTID="0" UPDATEDATE="2011-09-14" STARTDATE="1900-01-01" ENDDATE="2079-06-06" ISACTUAL="1" ISACTIVE="1"/>
<OBJECT ID="1234262" OBJECTID="996916" OBJECTGUID="632067b3-f39d-4f98-889a-547a58f4bccc" CHANGEID="2731136" NAME="Кленовая" TYPENAME="ул" LEVEL="8" OPERTYPEID="10" PREVID="0" NEXTID="1921662" UPDATEDATE="2021-08-19" STARTDATE="2018-09-25" ENDDATE="2021-08-19" ISACTUAL="0" ISACTIVE="0"/>
</ADDRESSOBJECTS>
`
	DataParams = `<?xml version="1.0" encoding="UTF-8"?>                                                                                                               teaman@MacBook-Pro-Tatana go-xml-parser % head -c 1000 AS_HOUSES_PARAMS_20230306_0260d3dc-9456-4cde-b970-2a6446414dd7.XML
<PARAMS>
<PARAM ID="331465750" OBJECTID="58820971" CHANGEID="87905117" CHANGEIDEND="87906330" TYPEID="7" VALUE="46643446126" UPDATEDATE="2019-07-10" STARTDATE="2018-05-28" ENDDATE="2019-01-14" />
<PARAM ID="331466052" OBJECTID="58820971" CHANGEID="87905117" CHANGEIDEND="87906537" TYPEID="6" VALUE="46243846007" UPDATEDATE="2019-07-10" STARTDATE="2018-05-28" ENDDATE="2019-02-03" />
<PARAM ID="331466104" OBJECTID="58820971" CHANGEID="87905117" CHANGEIDEND="87906537" TYPEID="13" VALUE="467490004960006000020074000000005" UPDATEDATE="2019-07-10" STARTDATE="2018-05-28" ENDDATE="2019-02-03" />
<PARAM ID="331468867" OBJECTID="58820971" CHANGEID="87906537" CHANGEIDEND="0" TYPEID="13" VALUE="467490004960006000020074000000000" UPDATEDATE="2019-07-10" STARTDATE="2019-02-03" ENDDATE="2079-06-06" />
</PARAMS>`
)
