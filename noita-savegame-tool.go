package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/rivo/tview"
)

// create a custom Noita world status struct
// for parcing the XML into

type noita_world struct {
	XMLName   xml.Name `xml:"Entity"`
	Text      string   `xml:",chardata"`
	Version   string   `xml:"_version,attr"`
	Name      string   `xml:"name,attr"`
	Serialize string   `xml:"serialize,attr"`
	Tags      string   `xml:"tags,attr"`
	Transform struct {
		Text      string `xml:",chardata"`
		PositionX string `xml:"position.x,attr"`
		PositionY string `xml:"position.y,attr"`
		Rotation  string `xml:"rotation,attr"`
		ScaleX    string `xml:"scale.x,attr"`
		ScaleY    string `xml:"scale.y,attr"`
	} `xml:"_Transform"`
	WorldStateComponent struct {
		Text                             string `xml:",chardata"`
		DEBUGLOADEDFROMAUTOSAVE          string `xml:"DEBUG_LOADED_FROM_AUTOSAVE,attr"`
		DEBUGLOADEDFROMOLDVERSION        string `xml:"DEBUG_LOADED_FROM_OLD_VERSION,attr"`
		ENDINGHAPPINESS                  string `xml:"ENDING_HAPPINESS,attr"`
		ENDINGHAPPINESSFRAMES            string `xml:"ENDING_HAPPINESS_FRAMES,attr"`
		ENDINGHAPPINESSHAPPENING         string `xml:"ENDING_HAPPINESS_HAPPENING,attr"`
		EVERYTHINGTOGOLD                 string `xml:"EVERYTHING_TO_GOLD,attr"`
		INFINITEGOLDHAPPENING            string `xml:"INFINITE_GOLD_HAPPENING,attr"`
		Enabled                          string `xml:"_enabled,attr"`
		Clouds01Target                   string `xml:"clouds_01_target,attr"`
		Clouds02Target                   string `xml:"clouds_02_target,attr"`
		ConsumeActions                   string `xml:"consume_actions,attr"`
		DamageFlashMultiplier            string `xml:"damage_flash_multiplier,attr"`
		DayCount                         string `xml:"day_count,attr"`
		Fog                              string `xml:"fog,attr"`
		FogTarget                        string `xml:"fog_target,attr"`
		GlobalGenomeRelationsModifier    string `xml:"global_genome_relations_modifier,attr"`
		GoreMultiplier                   string `xml:"gore_multiplier,attr"`
		GradientSkyAlphaTarget           string `xml:"gradient_sky_alpha_target,attr"`
		IntroWeather                     string `xml:"intro_weather,attr"`
		IsInitialized                    string `xml:"is_initialized,attr"`
		LightningCount                   string `xml:"lightning_count,attr"`
		MFlashAlpha                      string `xml:"mFlashAlpha,attr"`
		MaterialEverythingToGold         string `xml:"material_everything_to_gold,attr"`
		MaterialEverythingToGoldStatic   string `xml:"material_everything_to_gold_static,attr"`
		ModsHaveBeenActiveDuringThisRun  string `xml:"mods_have_been_active_during_this_run,attr"`
		NextCutThroughWorldID            string `xml:"next_cut_through_world_id,attr"`
		NextPortalID                     string `xml:"next_portal_id,attr"`
		OpenFogOfWarEverywhere           string `xml:"open_fog_of_war_everywhere,attr"`
		PerkGoldIsForever                string `xml:"perk_gold_is_forever,attr"`
		PerkHpDropChance                 string `xml:"perk_hp_drop_chance,attr"`
		PerkInfiniteSpells               string `xml:"perk_infinite_spells,attr"`
		PerkRatsPlayerFriendly           string `xml:"perk_rats_player_friendly,attr"`
		PerkTrickKillsBloodMoney         string `xml:"perk_trick_kills_blood_money,attr"`
		PlayerDidDamageOver1milj         string `xml:"player_did_damage_over_1milj,attr"`
		PlayerDidInfiniteSpellCount      string `xml:"player_did_infinite_spell_count,attr"`
		PlayerLivingWithMinusHp          string `xml:"player_living_with_minus_hp,attr"`
		PlayerPolymorphCount             string `xml:"player_polymorph_count,attr"`
		PlayerPolymorphRandomCount       string `xml:"player_polymorph_random_count,attr"`
		PlayerSpawnLocationX             string `xml:"player_spawn_location.x,attr"`
		PlayerSpawnLocationY             string `xml:"player_spawn_location.y,attr"`
		Rain                             string `xml:"rain,attr"`
		RainTarget                       string `xml:"rain_target,attr"`
		SessionStatFile                  string `xml:"session_stat_file,attr"`
		SkySunsetAlphaTarget             string `xml:"sky_sunset_alpha_target,attr"`
		Time                             string `xml:"time,attr"`
		TimeDt                           string `xml:"time_dt,attr"`
		TimeTotal                        string `xml:"time_total,attr"`
		TrickKillGoldMultiplier          string `xml:"trick_kill_gold_multiplier,attr"`
		TwitchHasBeenActiveDuringThisRun string `xml:"twitch_has_been_active_during_this_run,attr"`
		Wind                             string `xml:"wind,attr"`
		WindSpeed                        string `xml:"wind_speed,attr"`
		WindSpeedSin                     string `xml:"wind_speed_sin,attr"`
		WindSpeedSinT                    string `xml:"wind_speed_sin_t,attr"`
		LuaGlobals                       struct {
			Text string `xml:",chardata"`
			E    []struct {
				Text  string `xml:",chardata"`
				Key   string `xml:"key,attr"`
				Value string `xml:"value,attr"`
			} `xml:"E"`
		} `xml:"lua_globals"`
		PendingPortals      string `xml:"pending_portals"`
		ApparitionsPerLevel struct {
			Text      string `xml:",chardata"`
			Primitive string `xml:"primitive"`
		} `xml:"apparitions_per_level"`
		NpcParties       string `xml:"npc_parties"`
		OrbsFoundThisrun struct {
			Text      string   `xml:",chardata"`
			Primitive []string `xml:"primitive"`
		} `xml:"orbs_found_thisrun"`
		Flags struct {
			Text   string   `xml:",chardata"`
			String []string `xml:"string"`
		} `xml:"flags"`
		ChangedMaterials struct {
			Text   string   `xml:",chardata"`
			String []string `xml:"string"`
		} `xml:"changed_materials"`
		CutsThroughWorld string `xml:"cuts_through_world"`
	} `xml:"WorldStateComponent"`
	Entity struct {
		Text      string `xml:",chardata"`
		Version   string `xml:"_version,attr"`
		Name      string `xml:"name,attr"`
		Serialize string `xml:"serialize,attr"`
		Tags      string `xml:"tags,attr"`
		Transform struct {
			Text      string `xml:",chardata"`
			PositionX string `xml:"position.x,attr"`
			PositionY string `xml:"position.y,attr"`
			Rotation  string `xml:"rotation,attr"`
			ScaleX    string `xml:"scale.x,attr"`
			ScaleY    string `xml:"scale.y,attr"`
		} `xml:"_Transform"`
		PlayerStatsComponent struct {
			Text    string `xml:",chardata"`
			Enabled string `xml:"_enabled,attr"`
			Lives   string `xml:"lives,attr"`
			MaxHp   string `xml:"max_hp,attr"`
			Speed   string `xml:"speed,attr"`
		} `xml:"PlayerStatsComponent"`
	} `xml:"Entity"`
}

func main() {
	//read in the xml world state file
	var savefile string = getsavelocation()
	var infinatespellcount string

	xmlFile := readworldstate(savefile)
	// print is sucessfull
	fmt.Println("Sucessfully Opened World State")

	//convert xml to a byte array so we can unmarshal it
	worldstatexmlbytevalue, _ := ioutil.ReadAll(xmlFile)

	//initalize our world state array
	var noita_world noita_world

	//unmarshal our byte array with the struct we created
	xml.Unmarshal(worldstatexmlbytevalue, &noita_world)

	fmt.Println("Does player have the perk gold is forever? " + noita_world.WorldStateComponent.PerkGoldIsForever)
	fmt.Println("How many times hase player done damage over 1mill with a single spell? " + noita_world.WorldStateComponent.PlayerDidInfiniteSpellCount)
	fmt.Scanln(&infinatespellcount)
	noita_world.WorldStateComponent.PlayerDidInfiniteSpellCount = infinatespellcount
	fmt.Println("How many times hase player done damage over 1mill with a single spell? " + noita_world.WorldStateComponent.PlayerDidInfiniteSpellCount)

	// write out the world state xml file
	writeworldstate(savefile, xmlFile, noita_world)
}

func editAStat() {

}

func getsavelocation() string {
	//get the save file location, probably from steam?
	var savefile string
	savefile = "/mnt/c/Users/bmettenbrink3/Desktop/save00/world_state.xml"
	return savefile
}

func readworldstate(savefile string) *os.File {
	//read in the XML file
	xmlFile, err := os.Open(savefile)
	if err != nil {
		fmt.Println(err)
	}
	return xmlFile
}

func writeworldstate(savefile string, xmlFile *os.File, noita_world noita_world) {
	//write out XML file to disk
	file, _ := xml.MarshalIndent(noita_world, "", "")
	ioutil.WriteFile(savefile, file, 0644)
	xmlFile.Close()
}

func whichThingToEdit() {
	println("What would you like to edit?")
	println(args ...Type)
}

func readplayerstate() {

}

func writeplayerstate() {

}
