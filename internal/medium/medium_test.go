package medium

import (
	"strings"
	"testing"
)

const sampleFeed = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/">
<channel>
  <title>Stories on Medium</title>
  <item>
    <title><![CDATA[Getting Started with Terraform]]></title>
    <link>https://medium.com/@someone/getting-started-abc123?source=rss-xyz------2</link>
    <category><![CDATA[terraform]]></category>
    <category><![CDATA[devops]]></category>
    <pubDate>Wed, 15 Apr 2026 06:17:17 GMT</pubDate>
    <content:encoded><![CDATA[<figure><img alt="" src="https://cdn-images-1.medium.com/max/1024/cover.png" /></figure><h3>Intro</h3><p>Terraform lets you describe <strong>infrastructure</strong> as code &amp; apply it safely.</p><p>Second paragraph.</p>]]></content:encoded>
  </item>
</channel>
</rss>`

func TestParse(t *testing.T) {
	posts, err := Parse(strings.NewReader(sampleFeed))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(posts) != 1 {
		t.Fatalf("got %d posts, want 1", len(posts))
	}
	p := posts[0]

	checks := []struct{ name, got, want string }{
		{"title", p.Title, "Getting Started with Terraform"},
		{"url (query stripped)", p.URL, "https://medium.com/@someone/getting-started-abc123"},
		{"image", p.ImageURL, "https://cdn-images-1.medium.com/max/1024/cover.png"},
		{"summary (tags stripped, entities decoded)", p.Summary, "Terraform lets you describe infrastructure as code & apply it safely. Second paragraph."},
		{"tags", strings.Join(p.Tags, ","), "terraform,devops"},
		{"date", p.PublishedAt.Format("2006-01-02"), "2026-04-15"},
	}
	for _, c := range checks {
		if c.got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, c.got, c.want)
		}
	}
}

func TestSummarizeTruncatesAtWord(t *testing.T) {
	body := "<p>" + strings.Repeat("word ", 100) + "</p>"
	got := summarize(body, 50)
	if !strings.HasSuffix(got, "…") {
		t.Errorf("expected ellipsis, got %q", got)
	}
	if strings.HasSuffix(strings.TrimSuffix(got, "…"), " ") || strings.Contains(got, "wor…") {
		t.Errorf("expected a clean word cut, got %q", got)
	}
}
