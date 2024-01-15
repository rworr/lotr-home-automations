package lotr_gg_service

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// Fetch gear information from a given character page

const (
	baseUrlTemplate   = "https://lotr.gg%vgear/"
	gearBucketClass   = "bucket-item__main "
	gearNameClass     = "gear-item"
	gearDataAttr      = "data-bs-toggle"
	quantityNodeClass = "bucket-item__quantity"
)

type GearLevels []GearLevel

type GearLevel struct {
	Level string
	Gear  GearMap
}

type GearMap map[string]int

func GetCharacterGear(character string, characterUrls CharacterUrls) (GearLevels, error) {
	body, err := callWebsite(fmt.Sprintf(baseUrlTemplate, characterUrls[character]))
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	gearLevels := make(GearLevels, 0)
	crawlForGearTier(doc, &gearLevels)
	return gearLevels, nil
}

func crawlForGearTier(node *html.Node, gearLevels *GearLevels) {
	var gearLevel *GearLevel
	if node.Type == html.ElementNode && node.Data == "h2" &&
		node.FirstChild.Type == html.TextNode && strings.Contains(node.FirstChild.Data, "Gear Level") {
		println(node.FirstChild.Data)
		gearLevel = &GearLevel{
			Level: strings.TrimSpace(node.FirstChild.Data),
			Gear:  make(GearMap),
		}
	}

	if gearLevel != nil {
		// Gear tier is a dead-end node, need to iterate from parent level
		for child := node.Parent.FirstChild; child != nil; child = child.NextSibling {
			crawlForGear(child, gearLevel)
		}
		*gearLevels = append(*gearLevels, *gearLevel)
	} else {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawlForGearTier(child, gearLevels)
		}
	}
}

func crawlForGear(node *html.Node, gearLevel *GearLevel) {
	var containsGear bool
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == gearBucketClass {
			containsGear = true
		}
	}

	if containsGear {
		var gearName string
		var gearQuantity int
		for child := node.FirstChild; child != nil && (gearName == "" || gearQuantity == 0); child = child.NextSibling {
			foundName, foundQuantity := crawlForGearInfo(child)
			if foundName != "" {
				gearName = foundName
			}
			if foundQuantity != 0 {
				gearQuantity = foundQuantity
			}
		}

		if gearName != "" && gearQuantity > 0 {
			gearLevel.Gear[gearName] += gearQuantity
		}
	} else {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawlForGear(child, gearLevel)
		}
	}
}

func crawlForGearInfo(node *html.Node) (string, int) {
	var hasGearClass, hasTooltip bool
	var title string
	var quantity int

	for _, attr := range node.Attr {
		switch {
		case attr.Key == "class" && attr.Val == gearNameClass:
			hasGearClass = true
		case attr.Key == "class" && attr.Val == quantityNodeClass:
			clean := strings.TrimSpace(node.FirstChild.Data)
			quantity, _ = strconv.Atoi(clean)
		case attr.Key == gearDataAttr && attr.Val == "tooltip":
			hasTooltip = true
		case attr.Key == "title":
			title = attr.Val
		}
	}

	if !hasGearClass || !hasTooltip {
		title = ""
	}

	for child := node.FirstChild; child != nil && (title == "" || quantity == 0); child = child.NextSibling {
		foundTitle, foundQuantity := crawlForGearInfo(child)
		if foundTitle != "" {
			title = foundTitle
		}
		if foundQuantity != 0 {
			quantity = foundQuantity
		}
	}

	return title, quantity
}
