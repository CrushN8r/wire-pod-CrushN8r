package wirepod_ttr

import (
    "strings"
)

// The special Character Replacements
var specialCharactersReplacements = strings.NewReplacer(
    // quotes, apostrophes & misc
    "’", "'", 
    "‘", "'", 
    "“", "\"", 
    "”", "\"",
    "—", ", ", 
    "–", ", ", 
    "…", ",",  
    "...", ",",  
    "*", "",

    // bullets
    "•", "-  ", "‣", "-  ", "◦", "-  ",

    // fractions
    "¼", "1/4",  
    "½", "1/2",  
    "¾", "3/4",

    "1̅", "1 recurring",
    "2̅", "2 recurring",
    "3̅", "3 recurring",
    "4̅", "4 recurring",
    "5̅", "5 recurring",
    "6̅", "6 recurring",
    "7̅", "7 recurring",
    "8̅", "8 recurring",
    "9̅", "9 recurring",

    // URL encoded
    "%23|\\#", "hashtag",
    "%24|\\$", "Canadian Dollars, Ay?",
    "%26|\\&", "and",
    "%40|\\@", " at ",
    "\u00A0", " ", // &nbsp;

    // Mathematical and Related Symbols
    "±", "plus minus",  
    "÷", "divided by",  
    "√", "square root",  
    "∞", "infinity",  
    "≈", "almost equals",  
    "≠", "is not equal to",  
    "≡", "is equal to",  
    "≤", "is less than or equal to",  
    "≥", "is greater than or equal to",  
    "°", "degrees",  
    "π", "pi",  
    "∆", "delta",  
    "∑", "sum",  
    "∏", "product",  
    "×", "multiply by",  

    // Currency Symbols
    "€", "EUR",  
    "£", "GBP",  
    "¥", "JPY",  
    "₹", "INR", 
)

// phonetic map for Vector (my little experiment)
var phoneticReplacements = map[string]string{
    "CrushN8r": "Crush-Ehnaydir", 
    "flippin'": "fukkin", "flippin": "fukkin", "flipin": "fukkin", "frickin": "fukkin", "Boo-Ya": "Boo-Ya, Hawk, Too-Ya",
    "AI": "AY-EYE", "SpaceX": "Space-X", "Aha": "Ah,ha", "A-ha": "Ah,ha", "tastic": "tass-stick", "A-game": "AY-game", " A ": " AY ",
}

// Comprehensive emoji patterns (compiled once for performance)
// const emojiPattern = `[\x{1F600}-\x{1F64F}]|[\x{1F300}-\x{1F5FF}]|[\x{1F680}-\x{1F6FF}]|[\x{1F1E0}-\x{1F1FF}]|[\x{2600}-\x{26FF}]|[\x{2700}-\x{27BF}]|[\x{1F900}-\x{1F9FF}]|[\x{1F004}]|[\x{1F0CF}]|[\x{1F18E}]|[\x{1F191}-\x{1F251}]|[\x{2B50}]|[\x{1F9E6}-\x{1F9EB}]|[\x{1F914}]|[\x{2702}]|[\x{1F9C0}]|[\x{1F461}]|[\x{1F4AA}]|[\x{1F44B}]|[\x{1F48B}]|[\x{1F49C}]|[\x{1F9F1}]|[\x{1F9FB}]|[\x{1F525}]|[\x{2728}]|[\x{1F44F}]|[\x{1F389}]|[\x{1F60D}]|[\x{1F929}]|[\x{1F613}]|[\x{1F625}]|[\x{1F3C3}]|[\x{1F4A8}]|[\x{1F5A5}]|[\x{1F922}]|[\x{1F570}]|[\x{1F52E}]|[\x{1F950}]|[\x{1F9BB}]|[\x{1F94A}]|[\x{1F9D9}]|[\x{1F9F3}]|[\x{1FA78}]|[\x{1F57A}]|[\x{1F9CF}]|[\x{1F4B0}]|[\x{1FA99}]|[\x{1F9F8}]|[\x{1F9D8}]|[\x{1F9F7}]`

// Comprehensive emoji patterns (compiled once for performance)
const emojiPattern = `[\x{1F600}-\x{1F64F}]|` + // Emoticons
    `[\x{1F300}-\x{1F5FF}]|` + // Miscellaneous Symbols
    `[\x{1F680}-\x{1F6FF}]|` + // Transport and Map symbols
    `[\x{1F1E0}-\x{1F1FF}]|` + // Flags
    `[\x{2600}-\x{26FF}]|` +   // Miscellaneous Symbols
    `[\x{2700}-\x{27BF}]|` +   // Dingbats
    `[\x{1F900}-\x{1F9FF}]|` + // Supplemental Symbols
    `[\x{1F004}]|` +           // Mahjong Tile Red Dragon
    `[\x{1F0CF}]|` +           // Playing Card Black Joker
    `[\x{1F18E}]|` +           // Circled C
    `[\x{1F191}-\x{1F251}]|` + // Enclosed Alphanumeric Supplement
    `[\x{1F9E6}-\x{1F9EB}]|` + // Chess Symbols
    `[\x{1F914}]|` +           // Thinking Face
    `[\x{2702}]|` +            // Scissors
    `[\x{1F9C0}]|` +           // Cheese Wedge 
    `[\x{1F461}]|` +           // Woman's Shoe
    `[\x{1F4AA}]|` +           // Flexed Biceps
    `[\x{1F44B}]|` +           // Waving Hand
    `[\x{1F48B}]|` +           // Kiss Mark
    `[\x{1F49C}]|` +           // Orange Heart
    `[\x{1F9F1}]|` +           // Luggage
    `[\x{1F9FB}]|` +           // Haircut
    `[\x{1F525}]|` +           // Fire
    `[\x{2728}]|` +            // Sparkles
    `[\x{1F44F}]|` +           // Clapping Hands
    `[\x{1F389}]|` +           // Party Popper
    `[\x{1F60D}]|` +           // Smiling Face with Heart-Eyes
    `[\x{1F929}]|` +           // Star-Struck
    `[\x{1F613}]|` +           // Sweating Face
    `[\x{1F625}]|` +           // Disappointed Face
    `[\x{1F3C3}]|` +           // Person Running
    `[\x{1F4A8}]|` +           // Collision
    `[\x{1F5A5}]|` +           // Eye in Speech Bubble
    `[\x{1F922}]|` +           // Face with Hand Over Mouth
    `[\x{1F570}]|` +           // Old Key
    `[\x{1F52E}]|` +           // Crystal Ball
    `[\x{1F950}]|` +           // Croissant
    `[\x{1F9BB}]|` +           // Tamale
    `[\x{1F94A}]|` +           // Cooking
    `[\x{1F9D9}]|` +           // Mage
    `[\x{1F9F3}]|` +           // Glasses
    `[\x{1FA78}]|` +           // Gaming Device
    `[\x{1F57A}]|` +           // Man Juggling
    `[\x{1F9CF}]|` +           // Person in Lotus Position
    `[\x{1F4B0}]|` +           // Money Bag
    `[\x{1FA99}]|` +           // Rooting
    `[\x{1F9F7}]|` +           // Magic Wand
    `[\x{1F4A5}]`              // Collision
