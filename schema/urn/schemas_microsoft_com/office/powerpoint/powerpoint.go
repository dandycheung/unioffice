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

package powerpoint ;import (_f "encoding/xml";_a "fmt";_ed "github.com/unidoc/unioffice/v2";);func NewCT_Rel ()*CT_Rel {_ee :=&CT_Rel {};return _ee };func NewCT_Empty ()*CT_Empty {_edc :=&CT_Empty {};return _edc };type CT_Rel struct{

// Text Reference
IdAttr *string ;};

// Validate validates the Iscomment and its children
func (_ag *Iscomment )Validate ()error {return _ag .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};func (_gb *Iscomment )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _gb .CT_Empty .MarshalXML (e ,start );};

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_ad *CT_Rel )ValidateWithPath (path string )error {return nil };type Iscomment struct{CT_Empty };func (_cf *Iscomment )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_cf .CT_Empty =*NewCT_Empty ();for {_eeb ,_cg :=d .Token ();if _cg !=nil {return _a .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_cg );
};if _df ,_ff :=_eeb .(_f .EndElement );_ff &&_df .Name ==start .Name {break ;};};return nil ;};func (_adf *Textdata )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _adf .CT_Rel .MarshalXML (e ,start );};func (_ab *CT_Rel )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {if _ab .IdAttr !=nil {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0069\u0064"},Value :_a .Sprintf ("\u0025\u0076",*_ab .IdAttr )});
};e .EncodeToken (start );e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};func NewTextdata ()*Textdata {_ga :=&Textdata {};_ga .CT_Rel =*NewCT_Rel ();return _ga };func (_g *CT_Empty )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {for {_d ,_c :=d .Token ();
if _c !=nil {return _a .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_c );};if _da ,_dd :=_d .(_f .EndElement );_dd &&_da .Name ==start .Name {break ;};};return nil ;};func (_db *CT_Rel )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {for _ ,_cb :=range start .Attr {if _cb .Name .Local =="\u0069\u0064"{_be :=_cb .Value ;
_db .IdAttr =&_be ;continue ;};};for {_bde ,_dg :=d .Token ();if _dg !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_dg );};if _bb ,_dab :=_bde .(_f .EndElement );_dab &&_bb .Name ==start .Name {break ;
};};return nil ;};

// Validate validates the CT_Rel and its children
func (_bbe *CT_Rel )Validate ()error {return _bbe .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};type CT_Empty struct{};func NewIscomment ()*Iscomment {_bbf :=&Iscomment {};_bbf .CT_Empty =*NewCT_Empty ();return _bbf };type Textdata struct{CT_Rel };
func (_dda *Textdata )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_dda .CT_Rel =*NewCT_Rel ();for _ ,_fg :=range start .Attr {if _fg .Name .Local =="\u0069\u0064"{_dc :=_fg .Value ;_dda .IdAttr =&_dc ;continue ;};};for {_ce ,_acg :=d .Token ();
if _acg !=nil {return _a .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_acg );};if _fe ,_cbf :=_ce .(_f .EndElement );_cbf &&_fe .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_cgd *Textdata )ValidateWithPath (path string )error {if _eec :=_cgd .CT_Rel .ValidateWithPath (path );_eec !=nil {return _eec ;};return nil ;};func (_b *CT_Empty )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {e .EncodeToken (start );
e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};

// Validate validates the Textdata and its children
func (_fd *Textdata )Validate ()error {return _fd .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_ef *CT_Empty )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_fc *Iscomment )ValidateWithPath (path string )error {if _bbfb :=_fc .CT_Empty .ValidateWithPath (path );_bbfb !=nil {return _bbfb ;};return nil ;};

// Validate validates the CT_Empty and its children
func (_ac *CT_Empty )Validate ()error {return _ac .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};func init (){_ed .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );
_ed .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );
_ed .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );
_ed .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );
};