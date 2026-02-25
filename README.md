
## Pokedex REPL
This project is an interactive Pokédex REPL (Read–Eval–Print Loop) built for exploring the Pokémon world directly from your terminal.
With simple and intuitive commands, you can explore maps, discover Pokémon, catch them, and inspect their stats — all without leaving the command line.

## Features
- Explore Maps
  - `map` - View the next set of map locations
  - `mapb` - Go back to the previous set of map locations
  - `explore <location>` Discover pokemons available in a specific area
- Catch Pokemon
  - `catch <pokemon>` - Attempt to catch pokemon and add it to your collection
  - `inspect <pokemon>` – View detailed stats and information about a caught Pokémon
- Others
  - `help` - To check any other commands available for you
  - `exit` - If you are tired of catch and seing pokemons you can always come back later!

## Example of usage

```bash
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
...
Pokedex > explore eterna-city-area
Found Pokemon:
 - psyduck
 - golduck
 - magikarp
 - gyarados
 - barboach
 - whiscash
Pokedex > catch psyduck
Throwing a Pokeball at psyduck...
psyduck escaped!
Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  - hp: 20
  - attack: 10
  - defense: 55
  - special-attack: 15
  - special-defense: 20
  - speed: 80
Types:
  - water
Pokedex >  exit
Closing the Pokedex... Goodbye!
```

## Possibles Updates

- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon

