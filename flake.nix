{
    description = "Go devlopment environment flakes";

    inputs = {
      nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
      flake-utils.url = "github:numtide/flake-utils";
      templ.url = "github:a-h/templ";
    };

    outputs = { self, nixpkgs, flake-utils, templ }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gotools
            gopls
            delve
            golangci-lint
            templ.packages.${system}.templ
          ];

          shellHook = ''
            mkdir -p ~/.local/bin ~/go/bin
            ln -sf "$(command -v gopls)" ~/.local/bin/gopls
            ln -sf "$(command -v gopls)" ~/go/bin/gopls
            echo "Bin files and symlink created"

            cleanup() {
                rm -rf ~/.local/bin ~/go/bin 
                echo "Bin folder removed"
            }

            trap cleanup EXIT
          '';
        };
      });
}