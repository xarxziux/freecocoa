#!/usr/bin/env node

import NReadLines from 'n-readlines'

const headerRX = /^\[terrain_(?<header>\w*)\]/
const nameRX = /name *= *_\("(?:\?name:)?(?<name>[\w| |.]*)"\)/
const classRX = /class *= "(?<class>[\w| ]*)"/
const defenseRX = /defense_bonus *= (?<defense>\d*)/
const flagsRX = /^flags *=/
const moreFlagsRX = /^ *"/
const helpTextRX = /helptext/

const main = () => {
  const lines = new NReadLines('terrain.ruleset')
  let line = lines.next()

  console.log(`package stats

// AUTO-GENERATED FILE - DO NOT EDIT
// USE THE SCRIPT parse_terrain_ruleset.js TO RE-GENERATE IF NEEDED

import (
\t"github.com/xarxziux/freecocoa/src/models"
)

// TerrainStats lists all units combat stats as extracted from units.ruleset.
var TerrainStats = map[string]models.TerrainType{`)

  while (line) {
    const header = headerRX.exec(line)

    if (header !== null) {
      console.log(`\t"${header.groups.header}": {`)
      readClassSection(lines)
    }

    line = lines.next()
  }

  console.log('}')
}

const readClassSection = (lines) => {
  let line = lines.next()

  while (true) {
    const name = nameRX.exec(line)
    if (name !== null) {
      console.log(`\t\tName: "${name.groups.name}",`)
      line = lines.next()
    }

    const terrainClass = classRX.exec(line)
    if (terrainClass !== null) {
      console.log(`\t\tClass: models.${terrainClass.groups.class}Class,`)
      line = lines.next()
    }

    const defense = defenseRX.exec(line)
    if (defense !== null) {
      console.log(`\t\tDefenseBonus: ${defense.groups.defense},`)
      line = lines.next()
    }

    const flags = flagsRX.exec(line)
    if (flags !== null) {
      readTerrainFlags(lines, line)
      console.log('\t},')
      return
    }

    const helpText = helpTextRX.exec(line)
    if (helpText !== null) {
      console.log('\t},')
      return
    }

    line = lines.next()
  }
}

const readTerrainFlags = (lines, line) => {
  let nextLine = lines.next()
  let moreFlags = moreFlagsRX.exec(nextLine)

  while (moreFlags) {
    line = line + nextLine
    nextLine = lines.next()
    moreFlags = moreFlagsRX.exec(nextLine)
  }

  if (line.includes('NoCities')) {
    console.log('\t\tNoCities: true,')
  }

  if (line.includes('CanHaveRiver')) {
    console.log('\t\tCanHaveRiver: true,')
  }

  if (line.includes('UnsafeCoast')) {
    console.log('\t\tUnsafeCoast: true,')
  }

  if (line.includes('NoFortify')) {
    console.log('\t\tNoFortify: true,')
  }
}

main()
