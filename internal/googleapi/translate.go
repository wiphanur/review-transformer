package googleapi

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func TranslateText(targetLanguage, text string) (string, string, error) {
	// Support laguage text := "こんにちは世界"
	// Not Support language text := "Öncelikle evin konumu çok güzeldi. RER A istasyonuna 5 dk. Hem Disneyland’a hem Paris merkeze ulaşımı oldukça kolaydı. Monique geç kalmamıza rağmen bizi çok güzel karşıladı. Muhit güvenli ve sakindi. Etrafta marketler ve restaurantlar vardı. Ev kullanışlı sade ve rahattı. Monique’ e çok teşekkür ediyoruz."
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", "", fmt.Errorf("language.Parse: %w", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", "", fmt.Errorf("translate: %w", err)
	}
	if len(resp) == 0 {
		return "", "", fmt.Errorf("translate returned empty response to text: %s", text)
	}

	log.Printf("response from Google Translate API: %s", resp)
	return resp[0].Text, resp[0].Source.String(), nil
}

func DetectLanguage(text string) (*translate.Detection, error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("translate.NewClient: %w", err)
	}
	defer client.Close()
	lang, err := client.DetectLanguage(ctx, []string{text})
	if err != nil {
		return nil, fmt.Errorf("DetectLanguage: %w", err)
	}
	if len(lang) == 0 || len(lang[0]) == 0 {
		return nil, fmt.Errorf("DetectLanguage return value empty")
	}

	log.Printf("response from Google Trasnslate API: %v", lang)
	return &lang[0][0], nil
}
