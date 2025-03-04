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

package tempstorage ;import _gc "io";

// File is a representation of a storage file
// with Read, Write, Close and Name methods identical to os.File.
type File interface{_gc .Reader ;_gc .ReaderAt ;_gc .Writer ;_gc .Closer ;Name ()string ;};

// TempDir creates a name for a new temp directory using a pattern argument.
func TempDir (pattern string )(string ,error ){return _c .TempDir (pattern )};

// Add reads a file from a disk and adds it to the storage.
func Add (path string )error {return _c .Add (path )};

// TempFile creates new empty file in the storage and returns it.
func TempFile (dir ,pattern string )(File ,error ){return _c .TempFile (dir ,pattern )};

// RemoveAll removes all files according to the dir argument prefix.
func RemoveAll (dir string )error {return _c .RemoveAll (dir )};

// Open returns tempstorage File object by name.
func Open (path string )(File ,error ){return _c .Open (path )};var _c storage ;

// SetAsStorage changes temporary storage to newStorage.
func SetAsStorage (newStorage storage ){_c =newStorage };type storage interface{Open (_f string )(File ,error );TempFile (_b ,_a string )(File ,error );TempDir (_be string )(string ,error );RemoveAll (_fc string )error ;Add (_e string )error ;};