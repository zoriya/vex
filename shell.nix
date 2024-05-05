{pkgs ? import <nixpkgs> {}}:
  pkgs.mkShell {
    packages = with pkgs; [
      go
      wgo
      pgformatter
      postgresql
    ];
  }
