package migrations

import (
	"context"
	"log"
	"os"
	"review-transformer/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Language represents the structure for our languages
type Language struct {
	Code string `bson:"language_code"`
	Name string `bson:"language_name"`
}

func InsertSupportLanguages() error {
	collection := db.MongoClient.Database(os.Getenv("MONGO_DATABASE")).Collection("supportedLanguages")

	languages := []interface{}{
		Language{Code: "ar", Name: "Arabic"},
		Language{Code: "zh", Name: "Chinese (Simplified)"},
		Language{Code: "zh-Hant", Name: "Chinese (Traditional)"},
		Language{Code: "nl", Name: "Dutch"},
		Language{Code: "en", Name: "English"},
		Language{Code: "fr", Name: "French"},
		Language{Code: "de", Name: "German"},
		Language{Code: "id", Name: "Indonesian"},
		Language{Code: "it", Name: "Italian"},
		Language{Code: "ja", Name: "Japanese"},
		Language{Code: "ko", Name: "Korean"},
		Language{Code: "pt", Name: "Portuguese (Brazilian & Continental)"},
		Language{Code: "es", Name: "Spanish"},
		Language{Code: "th", Name: "Thai"},
		Language{Code: "tr", Name: "Turkish"},
		Language{Code: "vi", Name: "Vietnamese"},
	}

	for _, language := range languages {
		filter := bson.D{
			{Key: "language_code", Value: language.(Language).Code},
			{Key: "language_name", Value: language.(Language).Name},
		}
		update := bson.D{{Key: "$set", Value: bson.D{
			{Key: "language_code", Value: language.(Language).Code},
			{Key: "language_name", Value: language.(Language).Name},
		}}}

		opts := options.Update().SetUpsert(true)
		_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Inserted Supperted Language documents")
	return nil
}

func InsertLanguageCodes() error {
	collection := db.MongoClient.Database(os.Getenv("MONGO_DATABASE")).Collection("languages")

	languages := []interface{}{
		Language{"alpha2", "English"},
		Language{"aa", "Afar"},
		Language{"ab", "Abkhazian"},
		Language{"ae", "Avestan"},
		Language{"af", "Afrikaans"},
		Language{"ak", "Akan"},
		Language{"am", "Amharic"},
		Language{"an", "Aragonese"},
		Language{"ar", "Arabic"},
		Language{"as", "Assamese"},
		Language{"av", "Avaric"},
		Language{"ay", "Aymara"},
		Language{"az", "Azerbaijani"},
		Language{"ba", "Bashkir"},
		Language{"be", "Belarusian"},
		Language{"bg", "Bulgarian"},
		Language{"bh", "Bihari languages"},
		Language{"bi", "Bislama"},
		Language{"bm", "Bambara"},
		Language{"bn", "Bengali"},
		Language{"bo", "Tibetan"},
		Language{"br", "Breton"},
		Language{"bs", "Bosnian"},
		Language{"ca", "Catalan; Valencian"},
		Language{"ce", "Chechen"},
		Language{"ch", "Chamorro"},
		Language{"co", "Corsican"},
		Language{"cr", "Cree"},
		Language{"cs", "Czech"},
		Language{"cu", "Church Slavic; Old Slavonic; Church Slavonic; Old Bulgarian; Old Church Slavonic"},
		Language{"cv", "Chuvash"},
		Language{"cy", "Welsh"},
		Language{"da", "Danish"},
		Language{"de", "German"},
		Language{"dv", "Divehi; Dhivehi; Maldivian"},
		Language{"dz", "Dzongkha"},
		Language{"ee", "Ewe"},
		Language{"el", "Greek, Modern (1453-)"},
		Language{"en", "English"},
		Language{"eo", "Esperanto"},
		Language{"es", "Spanish; Castilian"},
		Language{"et", "Estonian"},
		Language{"eu", "Basque"},
		Language{"fa", "Persian"},
		Language{"ff", "Fulah"},
		Language{"fi", "Finnish"},
		Language{"fj", "Fijian"},
		Language{"fo", "Faroese"},
		Language{"fr", "French"},
		Language{"fy", "Western Frisian"},
		Language{"ga", "Irish"},
		Language{"gd", "Gaelic; Scottish Gaelic"},
		Language{"gl", "Galician"},
		Language{"gn", "Guarani"},
		Language{"gu", "Gujarati"},
		Language{"gv", "Manx"},
		Language{"ha", "Hausa"},
		Language{"he", "Hebrew"},
		Language{"hi", "Hindi"},
		Language{"ho", "Hiri Motu"},
		Language{"hr", "Croatian"},
		Language{"ht", "Haitian; Haitian Creole"},
		Language{"hu", "Hungarian"},
		Language{"hy", "Armenian"},
		Language{"hz", "Herero"},
		Language{"ia", "Interlingua (International Auxiliary Language Association)"},
		Language{"id", "Indonesian"},
		Language{"ie", "Interlingue; Occidental"},
		Language{"ig", "Igbo"},
		Language{"ii", "Sichuan Yi; Nuosu"},
		Language{"ik", "Inupiaq"},
		Language{"io", "Ido"},
		Language{"is", "Icelandic"},
		Language{"it", "Italian"},
		Language{"iu", "Inuktitut"},
		Language{"ja", "Japanese"},
		Language{"jv", "Javanese"},
		Language{"ka", "Georgian"},
		Language{"kg", "Kongo"},
		Language{"ki", "Kikuyu; Gikuyu"},
		Language{"kj", "Kuanyama; Kwanyama"},
		Language{"kk", "Kazakh"},
		Language{"kl", "Kalaallisut; Greenlandic"},
		Language{"km", "Central Khmer"},
		Language{"kn", "Kannada"},
		Language{"ko", "Korean"},
		Language{"kr", "Kanuri"},
		Language{"ks", "Kashmiri"},
		Language{"ku", "Kurdish"},
		Language{"kv", "Komi"},
		Language{"kw", "Cornish"},
		Language{"ky", "Kirghiz; Kyrgyz"},
		Language{"la", "Latin"},
		Language{"lb", "Luxembourgish; Letzeburgesch"},
		Language{"lg", "Ganda"},
		Language{"li", "Limburgan; Limburger; Limburgish"},
		Language{"ln", "Lingala"},
		Language{"lo", "Lao"},
		Language{"lt", "Lithuanian"},
		Language{"lu", "Luba-Katanga"},
		Language{"lv", "Latvian"},
		Language{"mg", "Malagasy"},
		Language{"mh", "Marshallese"},
		Language{"mi", "Maori"},
		Language{"mk", "Macedonian"},
		Language{"ml", "Malayalam"},
		Language{"mn", "Mongolian"},
		Language{"mr", "Marathi"},
		Language{"ms", "Malay"},
		Language{"mt", "Maltese"},
		Language{"my", "Burmese"},
		Language{"na", "Nauru"},
		Language{"nb", "Bokmål, Norwegian; Norwegian Bokmål"},
		Language{"nd", "Ndebele, North; North Ndebele"},
		Language{"ne", "Nepali"},
		Language{"ng", "Ndonga"},
		Language{"nl", "Dutch; Flemish"},
		Language{"nn", "Norwegian Nynorsk; Nynorsk, Norwegian"},
		Language{"no", "Norwegian"},
		Language{"nr", "Ndebele, South; South Ndebele"},
		Language{"nv", "Navajo; Navaho"},
		Language{"ny", "Chichewa; Chewa; Nyanja"},
		Language{"oc", "Occitan (post 1500); Provençal"},
		Language{"oj", "Ojibwa"},
		Language{"om", "Oromo"},
		Language{"or", "Oriya"},
		Language{"os", "Ossetian; Ossetic"},
		Language{"pa", "Panjabi; Punjabi"},
		Language{"pi", "Pali"},
		Language{"pl", "Polish"},
		Language{"ps", "Pushto; Pashto"},
		Language{"pt", "Portuguese"},
		Language{"qu", "Quechua"},
		Language{"rm", "Romansh"},
		Language{"rn", "Rundi"},
		Language{"ro", "Romanian; Moldavian; Moldovan"},
		Language{"ru", "Russian"},
		Language{"rw", "Kinyarwanda"},
		Language{"sa", "Sanskrit"},
		Language{"sc", "Sardinian"},
		Language{"sd", "Sindhi"},
		Language{"se", "Northern Sami"},
		Language{"sg", "Sango"},
		Language{"si", "Sinhala; Sinhalese"},
		Language{"sk", "Slovak"},
		Language{"sl", "Slovenian"},
		Language{"sm", "Samoan"},
		Language{"sn", "Shona"},
		Language{"so", "Somali"},
		Language{"sq", "Albanian"},
		Language{"sr", "Serbian"},
		Language{"ss", "Swati"},
		Language{"st", "Sotho, Southern"},
		Language{"su", "Sundanese"},
		Language{"sv", "Swedish"},
		Language{"sw", "Swahili"},
		Language{"ta", "Tamil"},
		Language{"te", "Telugu"},
		Language{"tg", "Tajik"},
		Language{"th", "Thai"},
		Language{"ti", "Tigrinya"},
		Language{"tk", "Turkmen"},
		Language{"tl", "Tagalog"},
		Language{"tn", "Tswana"},
		Language{"to", "Tonga (Tonga Islands)"},
		Language{"tr", "Turkish"},
		Language{"ts", "Tsonga"},
		Language{"tt", "Tatar"},
		Language{"tw", "Twi"},
		Language{"ty", "Tahitian"},
		Language{"ug", "Uighur; Uyghur"},
		Language{"uk", "Ukrainian"},
		Language{"ur", "Urdu"},
		Language{"uz", "Uzbek"},
		Language{"ve", "Venda"},
		Language{"vi", "Vietnamese"},
		Language{"vo", "Volapük"},
		Language{"wa", "Walloon"},
		Language{"wo", "Wolof"},
		Language{"xh", "Xhosa"},
		Language{"yi", "Yiddish"},
		Language{"yo", "Yoruba"},
		Language{"za", "Zhuang; Chuang"},
		Language{"zh", "Chinese"},
		Language{"zu", "Zulu"},
	}

	for _, language := range languages {
		filter := bson.D{
			{Key: "language_code", Value: language.(Language).Code},
			{Key: "language_name", Value: language.(Language).Name},
		}
		update := bson.D{{Key: "$set", Value: bson.D{
			{Key: "language_code", Value: language.(Language).Code},
			{Key: "language_name", Value: language.(Language).Name},
		}}}

		opts := options.Update().SetUpsert(true)
		_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Inserted Language Codes documents")
	return nil
}
