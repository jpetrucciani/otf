{ pkgs ? import
    (fetchTarball {
      name = "jpetrucciani-2023-12-26";
      url = "https://github.com/jpetrucciani/nix/archive/f3fb52047050d95a4a1da11a6ee40e03adc311e4.tar.gz";
      sha256 = "00iprklv9wxgra4ij9947njjgbmrks2hgq25y0s4ym3jrx0wrw51";
    })
    { }
}:
let
  name = "otf";

  tools = with pkgs; {
    cli = [
      coreutils
      nixpkgs-fmt
      nodejs_20
    ];
    go = [
      go
      go-tools
      gopls
    ];
    scripts = [ ];
  };

  paths = pkgs.lib.flatten [ (builtins.attrValues tools) ];
  env = pkgs.buildEnv {
    inherit name paths;
    buildInputs = paths;
  };
in
env.overrideAttrs (_: {
  NIXUP = "0.0.5";
})
