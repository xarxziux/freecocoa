#!/usr/bin/env node

import {createInterface} from 'readline'
import  {createReadStream} from 'fs'

const headerRX = /\[terrain_(?<terrain>.*)\]/
const nameRX = /name *= *_\("(?:\?name:)?(?<name>[\w| |\.]*)"\)/
const classRX = /class *= "(?<class>[\w| ]*)"/
const defenseRX = /defense_bonus *= (?<defense>\d*)/
const helpTextRX = /helptext/

//const hpRX = /hitpoints *= (?<hp>\d*)/
//const fpRX = /firepower *= (?<fp>\d*)/
//const cbRX = /CityBuster/

const main = () => {
    const file = createInterface({
        input: createReadStream('terrain.ruleset'),
        output: process.stdout,
        terminal: false
    });

    file.on('line', (line) => {
        parseLine(line)
    });
}

const parseLine = (line) => {
    const header = headerRX.exec(line)
    if (header !== null) {
        console.log(`\t"${header.groups.terrain}": {`)
        return
    }

    const name = nameRX.exec(line)
    if (name !== null) {
        console.log(`\t\tName: "${name.groups.name}",`)
        return
    }

    const terrainClass = classRX.exec(line)
    if (terrainClass !== null) {
        console.log(`\t\tClass: models.${terrainClass.groups.class}Class,`)
        return
    }

    const defense = defenseRX.exec(line)
    if (defense !== null) {
        console.log(`\t\tDefenseBonus: ${defense.groups.defense},`)
        return
    }

    const helpText = helpTextRX.exec(line)
    if (helpText !== null) {
        console.log(`\t},`)
        return
    }
}

main()
