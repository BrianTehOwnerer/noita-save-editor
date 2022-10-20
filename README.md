# Noita Save game editor - Work in progress
## needs to grab
### In game Info
- seed
  - in stats/session/newest  world_stats.xml
- time
  - same as seed
- wands
- spells
- items
- Perks
  - player.xml
- location
  - world_stats
- Orbs
  - player.xml
- NG+ status
- Fungal Shifts
  - world_state.xml
- Parallel Worlds
- Current Boss's killed
  - world_state.xml
### meta
- unlocks
  - world_state.xml
- achievements
- pillars
- secrets
  -
- enemies killed
  - https://noita.fandom.com/wiki/Enemies
## nice to have
- import wands?
- backup saves easily?
- tell next fungal shifts
- tell any interesting info about seed
- starting spells/items
-

# Process for loading info
1) grab savegame location
2) load file into program (not ram as it can get massive)
3) parse file and place important info into program
4) display data to user
5) allow edits to data
6) write data back


# XML GOLANG STRUCT TOOL
https://www.onlinetool.io/xmltogo/
