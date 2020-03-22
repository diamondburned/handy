{ pkgs ? import <nixpkgs> {} }:

pkgs.stdenv.mkDerivation rec {
	name = "gohandy";

	buildInputs = with pkgs; [
		gnome3.glib gnome3.gtk libhandy
		pkgconfig go
	];
}
