{ pkgs ? import
    (fetchTarball {
      name = "jpetrucciani-2024-09-02";
      url = "https://github.com/jpetrucciani/nix/archive/6c81869b06de00477f4c350d46c007e14fff21ff.tar.gz";
      sha256 = "1s8kp53hnv9zf61xsx6b9n5vm7iy74khj1dx8293y6282zrkihhg";
    })
    { }
}:
let
  name = "otf";

  tools = with pkgs; {
    cli = [
      coreutils
      nixpkgs-fmt
      nodejs_22
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
