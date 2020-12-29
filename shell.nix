{ pkgs ? import <nixpkgs> {} }:

let libhandy = pkgs.libhandy.overrideAttrs(old: {
	name = "libhandy-1.0.3";
	src  = builtins.fetchGit {
		url = "https://gitlab.gnome.org/GNOME/libhandy.git";
		rev = "7126d2e8da9dcdeb5751e60ff6a74930804f9637";
		ref = "libhandy-1-0";
	};
	patches = [];

	buildInputs = old.buildInputs ++ (with pkgs; [
		gnome3.librsvg
		gdk-pixbuf
	]);
});

in pkgs.stdenv.mkDerivation rec {
	name = "handy";

	buildInputs = [ libhandy ] ++ (with pkgs; [
		gnome3.glib gnome3.gtk
		pkgconfig go
	]);
}
