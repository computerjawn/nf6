{
  cfgFile ? "$HOME/.config/nf6-api-dev/config.yaml",
  dataDir ? "$HOME/.local/share/nf6-api-dev",
  go,
  writeShellApplication,
}:

writeShellApplication {
  name = "dev-server-api";
  runtimeInputs = [ go ];
  text = ''
    go run ./server-api/*.go --config "${cfgFile}" --dataDir "${dataDir}" "$@"
  '';
}
