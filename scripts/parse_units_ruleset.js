#!/usr/bin/env node

import nReadlines from "n-readlines"

const unitClassRX = /^\[\??unitclass_(?<unitClass>\w*)\]/
const flagsRX = /^flags *=/
const moreFlagsRX = /^ *\"/
const stringsValuesRX = /\"\w*\"/
const unitRX = /^\[\??unit_(?<unit>\w*)\]/
const nameRX = /^name *= *_\("(?:\?unit:)?(?<name>[\w| |\.]*)"\)/
const classRX = /^class *= "(?<class>[\w| ]*)"/
const attackRX = /^attack *= (?<attack>\d*)/
const defenseRX = /^defense *= (?<defense>\d*)/
const hpRX = /^hitpoints *= (?<hp>\d*)/
const fpRX = /^firepower *= (?<fp>\d*)/
const helpTextRX = /^helptext/

const main = () => {
    const lines = new nReadlines("units.ruleset")
    let line = lines.next()

    console.log(`package stats

// AUTO-GENERATED FILE - DO NOT EDIT
// USE THE SCRIPT parse_units_ruleset.js TO RE-GENERATE IF NEEDED

import (
	"github.com/xarxziux/freecocoa/src/models"
)

// UnitClassStats lists all unit classes and combat-relevant flags as extracted
// from units.ruleset.
var UnitClassStats = map[string]models.UnitClass{`)

    while (unitRX.exec(line) === null) {
        const unitClass = unitClassRX.exec(line)
        if (unitClass !== null) {
            console.log(`\t"${unitClass.groups.unitClass}": {`)
            readClassSection(lines)
            console.log("\t},")
        }

        line = lines.next()
    }

    console.log("}")
    console.log()
    console.log(`// UnitStats lists all units combat stats and combat-relevant flags as extracted
// from units.ruleset.
var UnitStats = map[string]models.UnitDetails{`)

    while (line) {
        const unit = unitRX.exec(line)
        if (unit !== null) {
            console.log(`\t"${unit.groups.unit}": {`)
            readUnitsSection(lines)
        }

        line = lines.next()
    }

    console.log("}")
}

const readClassSection = (lines) => {
    return
}

const readUnitsSection = (lines) => {
    let line = lines.next()

    while (true) {
        const name = nameRX.exec(line)
        if (name !== null) {
            console.log(`\t\tName: "${name.groups.name}",`)
        }

        const unitClassName = classRX.exec(line)
        if (unitClassName !== null) {
            console.log(`\t\tClass: UnitClassStats["${unitClassName.groups.class.replace(/ /g, "")}"],`)
        }

        const attack = attackRX.exec(line)
        if (attack !== null) {
            console.log(`\t\tAttack: ${attack.groups.attack},`)
        }

        const defense = defenseRX.exec(line)
        if (defense !== null) {
            console.log(`\t\tDefense: ${defense.groups.defense},`)
        }

        const hp = hpRX.exec(line)
        if (hp !== null) {
            console.log(`\t\tHP: ${hp.groups.hp},`)
        }

        const fp = fpRX.exec(line)
        if (fp !== null) {
            console.log(`\t\tFP: ${fp.groups.fp},`)
        }

        const flags = flagsRX.exec(line)
        if (flags !== null) {
            readUnitFlags(lines, line)
        }

        const helpText = helpTextRX.exec(line)
        if (helpText !== null) {
            console.log(`\t},`)
            return
        }

        line = lines.next()
    }
}

const readUnitFlags = (lines, line) => {
    while (true) {
        const nextLine = lines.next()

        const moreFlags = moreFlagsRX.exec(nextLine)
        if (moreFlags ==  null) {
            break
        }

        line = line + nextLine
    }

    if (line.includes("CityBuster")) {
        console.log ("\t\tCityBuster: true,")
    }

    if (line.includes("AirAttacker")) {
        console.log ("\t\tAirAttacker: true,")
    }

    if (line.includes("Horse")) {
        console.log ("\t\tHorse: true,")
    }

    if (line.includes("Submarine")) {
        console.log ("\t\tSubmarine: true,")
    }

    if (line.includes("BadCityDefender")) {
        console.log ("\t\tBadCityDefender: true,")
    }

    if (line.includes("OnlyNativeAttack")) {
        console.log ("\t\tOnlyNativeAttack: true,")
    }
}

main()
