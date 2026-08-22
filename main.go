package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/xtls/xray-core/app/router"
	"google.golang.org/protobuf/proto"
)

type CategoryRule struct {
	Tag      string   
	BaseURLs []string
}

func opencckGroup(domain, group string) string {
	return fmt.Sprintf("https://%s/?format=text&group=%s", domain, group)
}

var categories = []CategoryRule{
	{
		Tag: "AI",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "ai"),
			"https://beta.iplist.opencck.org/?format=text&site=google%40aistudio&site=google%40gemini&site=google%40google-gemini&site=google%40notebooklm",
		},
	},
	{
		Tag: "ANIME",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "anime"),
			"https://beta.iplist.opencck.org/?format=text&site=aniliberty.top&site=mangalib.me",
		},
	},
	{
		Tag: "ART",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "art"),
			"https://beta.iplist.opencck.org/?format=text&site=wattpad.com",
		},
	},
	{
		Tag: "CASINO",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "casino"),
		},
	},
	{
		Tag: "DISCORD",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "discord"),
		},
	},
	{
		Tag: "EDUCATION",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "education"),
		},
	},
	{
		Tag: "FINANCE",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "finance"),
		},
	},
	{
		Tag: "GAMES",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "games"),
			"https://beta.iplist.opencck.org/?format=text&site=brawlstars.com&site=ea.com%40battlefield1&site=ea.com%40battlefield2042&site=ea.com%40cdn&site=leagueoflegends.com&site=centurygame.com&site=blitz.gg",
		},
	},
	{
		Tag: "HOSTING",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "hosting"),
			"https://beta.iplist.opencck.org/?format=text&site=cloudflare.com&site=digitalocean.com",
		},
	},
	{
		Tag: "MESSENGERS",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "messengers"),
			"https://beta.iplist.opencck.org/?format=text&site=imo.im&site=slack.com",
		},
	},
	{
		Tag: "MUSIC",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "music"),
		},
	},
	{
		Tag: "NEWS",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "news"),
		},
	},
	{
		Tag: "PORN",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "porn"),
		},
	},
	{
		Tag: "SHOP",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "shop"),
		},
	},
	{
		Tag: "SOCIALS",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "socials"),
			"https://beta.iplist.opencck.org/?format=text&site=bsky.app",
		},
	},
	{
		Tag: "TOOLS",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "tools"),
			"https://beta.iplist.opencck.org/?format=text&site=openwrt.org&site=speedtest.net&site=aidungeon.com&site=anydesk.com&site=cohere.com&site=figma.com&site=mega.io&site=themeforest.net&site=yootheme.com",
		},
	},
	{
		Tag: "TORRENT",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "torrent"),
		},
	},
	{
		Tag: "VIDEO",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "video"),
			"https://beta.iplist.opencck.org/?format=text&site=antifriz.tv",
		},
	},
	{
		Tag: "YOUTUBE",
		BaseURLs: []string{
			opencckGroup("iplist.opencck.org", "youtube"),
		},
	},
	{
		Tag: "TIKTOK", 
		BaseURLs: []string{
			"https://iplist.opencck.org/?format=text&site=tiktok.com",
		},
	},
	{
		Tag: "RUSSIA",
		BaseURLs: []string{
			opencckGroup("russia.iplist.opencck.org", "russia"),
		},
	},
	{
		Tag: "VK",
		BaseURLs: []string{
			opencckGroup("russia.iplist.opencck.org", "vk"),
		},
	},
	{
		Tag: "YANDEX",
		BaseURLs: []string{
			opencckGroup("russia.iplist.opencck.org", "yandex"),
		},
	},
}

var domainPattern = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

func main() {
	fmt.Println("=== Starting GeoData Builder (IPv4 Only) ===")

	client := &http.Client{Timeout: 40 * time.Second}

	var geoSiteEntries []*router.GeoSite
	var geoIPEntries []*router.GeoIP

	for _, cat := range categories {
		tag := strings.ToUpper(cat.Tag)
		fmt.Printf("\n[+] Processing Tag: %s\n", tag)

		uniqueDomains := make(map[string]bool)
		for _, baseURL := range cat.BaseURLs {
			domainURL := transformURL(baseURL, "domains")
			domains, err := fetchAndCleanDomains(client, domainURL)
			if err != nil {
				fmt.Printf("   [!] Skipped domain fetch (%s): %v\n", domainURL, err)
				continue
			}
			for _, d := range domains {
				uniqueDomains[d] = true
			}
			time.Sleep(300 * time.Millisecond)
		}

		if len(uniqueDomains) > 0 {
			var domainList []string
			for d := range uniqueDomains {
				domainList = append(domainList, d)
			}
			geoSiteEntries = append(geoSiteEntries, buildGeoSiteEntry(tag, domainList))
			fmt.Printf("   [✓] Geosite '%s': %d domains\n", tag, len(domainList))
		}

		uniqueCIDRs := make(map[string]bool)
		for _, baseURL := range cat.BaseURLs {
			cidr4URL := transformURL(baseURL, "cidr4")
			cidrs4, err := fetchLines(client, cidr4URL)
			if err != nil {
				fmt.Printf("   [!] Skipped CIDR fetch (%s): %v\n", cidr4URL, err)
				continue
			}
			for _, c := range cidrs4 {
				uniqueCIDRs[c] = true
			}
			time.Sleep(300 * time.Millisecond)
		}

		if len(uniqueCIDRs) > 0 {
			var cidrList []string
			for c := range uniqueCIDRs {
				cidrList = append(cidrList, c)
			}
			ipEntry, err := buildGeoIPEntry(tag, cidrList)
			if err == nil && len(ipEntry.Cidr) > 0 {
				geoIPEntries = append(geoIPEntries, ipEntry)
				fmt.Printf("   [✓] GeoIP '%s': %d IPv4 CIDR ranges\n", tag, len(ipEntry.Cidr))
			}
		}
	}

	if len(geoSiteEntries) > 0 {
		siteList := &router.GeoSiteList{Entry: geoSiteEntries}
		if err := saveProto("geosite.dat", siteList); err != nil {
			fmt.Printf("\nError saving geosite.dat: %v\n", err)
		} else {
			fmt.Println("\n[SUCCESS] geosite.dat successfully compiled!")
		}
	}

	if len(geoIPEntries) > 0 {
		ipList := &router.GeoIPList{Entry: geoIPEntries}
		if err := saveProto("geoip.dat", ipList); err != nil {
			fmt.Printf("Error saving geoip.dat: %v\n", err)
		} else {
			fmt.Println("[SUCCESS] geoip.dat successfully compiled!")
		}
	}
}

func transformURL(baseURL string, dataType string) string {
	u := baseURL
	u = strings.Replace(u, "format=json", "format=text", -1)

	if strings.Contains(u, "data=") {
		re := regexp.MustCompile(`data=[^&]+`)
		u = re.ReplaceAllString(u, "data="+dataType)
	} else {
		if strings.Contains(u, "?") {
			u += "&data=" + dataType
		} else {
			u += "?data=" + dataType
		}
	}
	return u
}

func fetchLines(client *http.Client, url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) Xray-GeoData-Builder/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status %d", resp.StatusCode)
	}

	var lines []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

func fetchAndCleanDomains(client *http.Client, url string) ([]string, error) {
	rawLines, err := fetchLines(client, url)
	if err != nil {
		return nil, err
	}

	var cleanDomains []string
	for _, line := range rawLines {
		if isValidXrayDomain(line) {
			cleanDomains = append(cleanDomains, line)
		}
	}
	return cleanDomains, nil
}

func isValidXrayDomain(line string) bool {
	if strings.HasPrefix(line, "full:") || strings.HasPrefix(line, "domain:") || strings.HasPrefix(line, "keyword:") {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			return false
		}
		return domainPattern.MatchString(parts[1])
	}

	if strings.HasPrefix(line, "regexp:") {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			return false
		}
		_, err := regexp.Compile(parts[1])
		return err == nil
	}

	return domainPattern.MatchString(line)
}

func buildGeoSiteEntry(tag string, lines []string) *router.GeoSite {
	geo := &router.GeoSite{
		CountryCode: tag,
		Domain:      make([]*router.Domain, 0, len(lines)),
	}

	for _, line := range lines {
		dType := router.Domain_Domain
		val := line

		if strings.HasPrefix(line, "full:") {
			dType = router.Domain_Full
			val = line[5:]
		} else if strings.HasPrefix(line, "domain:") {
			dType = router.Domain_Domain
			val = line[7:]
		} else if strings.HasPrefix(line, "regexp:") {
			dType = router.Domain_Regex
			val = line[7:]
		} else if strings.HasPrefix(line, "keyword:") {
			dType = router.Domain_Plain
			val = line[8:]
		}

		geo.Domain = append(geo.Domain, &router.Domain{
			Type:  dType,
			Value: val,
		})
	}
	return geo
}

func buildGeoIPEntry(tag string, lines []string) (*router.GeoIP, error) {
	geo := &router.GeoIP{
		CountryCode: tag,
		Cidr:        make([]*router.CIDR, 0, len(lines)),
	}

	for _, line := range lines {
		cidr, err := parseCIDR(line)
		if err != nil {
			continue 
		}
		geo.Cidr = append(geo.Cidr, cidr)
	}
	return geo, nil
}

func parseCIDR(line string) (*router.CIDR, error) {
	ip, ipnet, err := net.ParseCIDR(line)
	if err != nil {
		ip = net.ParseIP(line)
		if ip == nil {
			return nil, fmt.Errorf("invalid ip")
		}
		if ip.To4() != nil {
			_, ipnet, _ = net.ParseCIDR(line + "/32")
		} else {
			return nil, fmt.Errorf("ignoring ipv6")
		}
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return nil, fmt.Errorf("ignoring ipv6")
	}

	prefix, _ := ipnet.Mask.Size()
	return &router.CIDR{
		Ip:     ip4,
		Prefix: uint32(prefix),
	}, nil
}

func saveProto(path string, m proto.Message) error {
	data, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
