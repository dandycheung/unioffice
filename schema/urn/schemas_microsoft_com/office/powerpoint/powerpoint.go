//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package powerpoint ;import (_c "encoding/xml";_a "fmt";_f "github.com/unidoc/unioffice/v2";);

// Validate validates the CT_Rel and its children
func (_e *CT_Rel )Validate ()error {return _e .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};func (_gf *CT_Rel )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _gf .IdAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0069\u0064"},Value :_a .Sprintf ("\u0025\u0076",*_gf .IdAttr )});
};e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type Iscomment struct{CT_Empty };

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_gb *CT_Rel )ValidateWithPath (path string )error {return nil };

// Validate validates the Textdata and its children
func (_cf *Textdata )Validate ()error {return _cf .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};func NewTextdata ()*Textdata {_de :=&Textdata {};_de .CT_Rel =*NewCT_Rel ();return _de };

// Validate validates the CT_Empty and its children
func (_df *CT_Empty )Validate ()error {return _df .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};type CT_Rel struct{

// Text Reference
IdAttr *string ;};func (_aa *CT_Empty )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_bb ,_af :=d .Token ();if _af !=nil {return _a .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_af );
};if _fb ,_dd :=_bb .(_c .EndElement );_dd &&_fb .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_bgb *Iscomment )ValidateWithPath (path string )error {if _ad :=_bgb .CT_Empty .ValidateWithPath (path );_ad !=nil {return _ad ;};return nil ;};func NewCT_Rel ()*CT_Rel {_ab :=&CT_Rel {};return _ab };func (_fg *Textdata )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _fg .CT_Rel .MarshalXML (e ,start );};func (_ed *Textdata )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_ed .CT_Rel =*NewCT_Rel ();for _ ,_gd :=range start .Attr {if _gd .Name .Local =="\u0069\u0064"{_cdd :=_gd .Value ;
_ed .IdAttr =&_cdd ;continue ;};};for {_gca ,_cc :=d .Token ();if _cc !=nil {return _a .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_cc );};if _bac ,_eb :=_gca .(_c .EndElement );_eb &&_bac .Name ==start .Name {break ;
};};return nil ;};type Textdata struct{CT_Rel };

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_fd *CT_Empty )ValidateWithPath (path string )error {return nil };

// Validate validates the Iscomment and its children
func (_bc *Iscomment )Validate ()error {return _bc .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};func NewCT_Empty ()*CT_Empty {_d :=&CT_Empty {};return _d };func (_ddb *Iscomment )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _ddb .CT_Empty .MarshalXML (e ,start );};func NewIscomment ()*Iscomment {_abf :=&Iscomment {};_abf .CT_Empty =*NewCT_Empty ();return _abf };type CT_Empty struct{};func (_gfe *CT_Rel )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for _ ,_dfa :=range start .Attr {if _dfa .Name .Local =="\u0069\u0064"{_gg :=_dfa .Value ;
_gfe .IdAttr =&_gg ;continue ;};};for {_da ,_gc :=d .Token ();if _gc !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_gc );};if _bg ,_ge :=_da .(_c .EndElement );_ge &&_bg .Name ==start .Name {break ;
};};return nil ;};

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_caf *Textdata )ValidateWithPath (path string )error {if _deg :=_caf .CT_Rel .ValidateWithPath (path );_deg !=nil {return _deg ;};return nil ;};func (_b *CT_Empty )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {e .EncodeToken (start );
e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_ca *Iscomment )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_ca .CT_Empty =*NewCT_Empty ();for {_bgd ,_cb :=d .Token ();if _cb !=nil {return _a .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_cb );
};if _cd ,_ba :=_bgd .(_c .EndElement );_ba &&_cd .Name ==start .Name {break ;};};return nil ;};func init (){_f .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );
_f .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );
_f .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );
_f .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );
};