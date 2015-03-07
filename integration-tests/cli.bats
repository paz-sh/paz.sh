#!/usr/bin/env bats

load vars

@test "cli: show info" {
  run paz
  [ "$status" -eq 0 ]
  [[ ${lines[0]} =~ "NAME:" ]]
  [[ ${lines[1]} =~ "paz" ]]
}

@test "cli: show help" {
  run paz help
  [ "$status" -eq 0 ]
  [[ ${lines[0]} =~ "NAME:" ]]
  [[ ${lines[1]} =~ "paz" ]]
}

@test "cli: show error for unknown command" {
  run paz unknown-command
  [ "$status" -gt 0 ]
  [[ ${lines[0]} =~ "paz: unknown subcommand: \"unknown-command\"" ]]
  [[ ${lines[1]} =~ "paz help" ]]
}

@test "cli: show version" {
  run paz version
  [ "$status" -eq 0 ]
  [[ ${lines[0]} =~ "version" ]]
}

@test "cli: show provision error" {
  run paz provision
  [ "$status" -gt 0 ]
  [[ ${lines[0]} =~ "provision" ]]
}

@test "flag: show provision help" {
  run paz provision --help
  [ "$status" -gt 0 ]
  [[ ${lines[0]} =~ "Usage" ]]
}

@test "flag: show help" {
  run paz --help
  [ "$status" -gt 0 ]
  [[ ${lines[0]} =~ "Usage of paz:" ]]
}
