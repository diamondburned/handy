{ pkgs ? import <nixpkgs> {} }:

let updateScript = pkgs.runCommandLocal "notify" {
	src = ./.;
} ''
	gir_version() {
		[[ $(< "$1") =~ repository\ version=\"([0-9.]+)\" ]] && \
			echo -n "''${BASH_REMATCH[1]}"
	}

	main() {
		echo
		newHandy="${pkgs.libhandy.dev}/share/gir-1.0/Handy-1.gir"

		cmp -s "$src/Handy-1.gir" "$newHandy" && {
			echo "Local Handy-1.gir is the same as generated."
			return
		}

		echo "Local Handy-1.gir is older/different from upstream."
		echo "To synchronize them, run:"
		echo -e "\tcp -f $newHandy ./"
		echo -e "\tgo generate ./..."
	}

	main "$@"
	mkdir -p "$out"
'';

in pkgs.mkShell {
	buildInputs = with pkgs; [
		libhandy
		gnome3.glib
		gnome3.gtk
	];

	nativeBuildInputs = [
		updateScript
		pkgs.pkgconfig
		pkgs.go
	];
}
