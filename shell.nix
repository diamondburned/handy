{ pkgs ? import <nixpkgs> {} }:

let libhandy = pkgs.libhandy.overrideAttrs(old: {
	name = "libhandy-0.90.0";
	src  = builtins.fetchGit {
		url = "https://gitlab.gnome.org/GNOME/libhandy.git";
		rev = "c7aaf6f4f50b64ee55fcfee84000e9525fc5f93a";
	};
	patches = [];

	buildInputs = old.buildInputs ++ (with pkgs; [
		gnome3.librsvg
		gdk-pixbuf
	]);
});

in pkgs.stdenv.mkDerivation rec {
	name = "gohandy";

	buildInputs = [ libhandy ] ++ (with pkgs; [
		gnome3.glib gnome3.gtk
		pkgconfig go
	]);
}
