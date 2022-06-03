#!/usr/bin/env node

import {createInterface} from 'readline'
import  {createReadStream} from 'fs'

const unitRX = /\[\??unit_(?<unit>\w*)\]/
const nameRX = /name *= *_\("(?:\?unit:)?(?<name>[\w| ]*)"\)/
const classRX = /class *= "(?<class>[\w| ]*)"/
const attackRX = /attack *= (?<attack>\d*)/
const defenseRX = /defense *= (?<defense>\d*)/
const hpRX = /hitpoints *= (?<hp>\d*)/
const fpRX = /firepower *= (?<fp>\d*)/
const cbRX = /CityBuster/
const helpTextRX = /helptext/

const main = () => {
    const file = createInterface({
        input: createReadStream('units.ruleset'),
        output: process.stdout,
        terminal: false
    });

    console.log("Units := []UnitType{")

    file.on('line', (line) => {
        parseLine(line)
    });
}

const parseLine = (line) => {
    const header = unitRX.exec(line)
    if (header !== null) {
        console.log(`\t"${header.groups.unit}": {`)
        return
    }

    const name = nameRX.exec(line)
    if (name !== null) {
        console.log(`\t\tName: "${name.groups.name}",`)
        return
    }

    const unitClass = classRX.exec(line)
    if (unitClass !== null) {
        console.log(`\t\tClass: ${unitClass.groups.class.replace(/ /g, "")},`)
        return
    }

    const attack = attackRX.exec(line)
    if (attack !== null) {
        console.log(`\t\tAttack: ${attack.groups.attack},`)
        return
    }

    const defense = defenseRX.exec(line)
    if (defense !== null) {
        console.log(`\t\tDefense: ${defense.groups.defense},`)
        return
    }

    const hp = hpRX.exec(line)
    if (hp !== null) {
        console.log(`\t\tHP: ${hp.groups.hp},`)
        return
    }

    const fp = fpRX.exec(line)
    if (fp !== null) {
        console.log(`\t\tFP: ${fp.groups.fp},`)
        return
    }

    const cb = cbRX.exec(line)
    if (cb !== null) {
        console.log(`\t\tCityBuster: true,`)
        return
    }

    const helpText = helpTextRX.exec(line)
    if (helpText !== null) {
        console.log(`\t},`)
        return
    }
}

main()
