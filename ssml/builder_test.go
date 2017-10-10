package ssml

import (
	"testing"
	"time"

	"github.com/mikeflynn/go-alexa/ssml/amazoneffect"
	"github.com/mikeflynn/go-alexa/ssml/emphasis"
	"github.com/mikeflynn/go-alexa/ssml/pause"
)

func TestNewBuilder_ReturnsEmptySSML(t *testing.T) {
	b, err := NewBuilder()

	if err != nil {
		t.Fatalf("failed to get new builder: expected no error, got :%v", err)
	}

	actual := b.Build()
	expected := "<speak></speak>"
	if actual != expected {
		t.Errorf("output mismatch: expected %s, got %s", expected, actual)
	}
}

func TestBuilder_AppendPlainSpeech(t *testing.T) {
	b, _ := NewBuilder()

	b.AppendPlainSpeech("hello ")
	b.AppendPlainSpeech("world")

	actual := b.Build()
	expected := "<speak>hello world</speak>"
	if actual != expected {
		t.Errorf("output mismatch: expected %s, got %s", expected, actual)
	}
}

func TestBuilder_AppendAmazonEffect(t *testing.T) {
	tests := []struct {
		name     string
		effect   amazoneffect.Effect
		text     string
		expected string
	}{
		{
			name:     "whispered",
			effect:   amazoneffect.Whispered,
			text:     "test1",
			expected: `<speak><amazon:effect name="whispered">test1</amazon:effect><amazon:effect name="whispered">test1</amazon:effect></speak>`,
		},
		{
			name:     "custom",
			effect:   amazoneffect.Effect("custom"),
			text:     "test2",
			expected: `<speak><amazon:effect name="custom">test2</amazon:effect><amazon:effect name="custom">test2</amazon:effect></speak>`,
		},
	}

	for _, test := range tests {
		b, _ := NewBuilder()

		b.AppendAmazonEffect(test.effect, test.text)
		b.AppendAmazonEffect(test.effect, test.text)

		actual := b.Build()
		if actual != test.expected {
			t.Errorf("%s: output mismatch: expected %s, got %s", test.name, test.expected, actual)
		}
	}
}

func TestBuilder_AppendAudio(t *testing.T) {
	b, _ := NewBuilder()

	b.AppendAudio("source1")
	b.AppendAudio("source2")

	actual := b.Build()
	expected := `<speak><audio src="source1"/><audio src="source2"/></speak>`
	if actual != expected {
		t.Errorf("output mismatch: expected %s, got %s", expected, actual)
	}
}

func TestBuilder_AppendBreak(t *testing.T) {
	tests := []struct {
		name     string
		strength pause.Strength
		duration time.Duration
		expected string
	}{
		{
			name:     "default",
			strength: pause.Default,
			duration: time.Second,
			expected: `<speak><break strength="medium" time="1000ms"/><break strength="medium" time="1000ms"/></speak>`,
		},
		{
			name:     "none",
			strength: pause.None,
			duration: time.Second / 2,
			expected: `<speak><break strength="none" time="500ms"/><break strength="none" time="500ms"/></speak>`,
		},
		{
			name:     "x-weak",
			strength: pause.XWeak,
			duration: time.Second * 2,
			expected: `<speak><break strength="x-weak" time="2000ms"/><break strength="x-weak" time="2000ms"/></speak>`,
		},
		{
			name:     "weak",
			strength: pause.Weak,
			duration: time.Second * 3,
			expected: `<speak><break strength="weak" time="3000ms"/><break strength="weak" time="3000ms"/></speak>`,
		},
		{
			name:     "medium",
			strength: pause.Medium,
			duration: time.Second * 4,
			expected: `<speak><break strength="medium" time="4000ms"/><break strength="medium" time="4000ms"/></speak>`,
		},
		{
			name:     "strong",
			strength: pause.Strong,
			duration: time.Second * 5,
			expected: `<speak><break strength="strong" time="5000ms"/><break strength="strong" time="5000ms"/></speak>`,
		},
		{
			name:     "x-strong",
			strength: pause.XStrong,
			duration: time.Second * 6,
			expected: `<speak><break strength="x-strong" time="6000ms"/><break strength="x-strong" time="6000ms"/></speak>`,
		},
		{
			name:     "custom",
			strength: pause.Strength("custom"),
			duration: time.Second * 7,
			expected: `<speak><break strength="custom" time="7000ms"/><break strength="custom" time="7000ms"/></speak>`,
		},
	}

	for _, test := range tests {
		b, _ := NewBuilder()

		b.AppendBreak(test.strength, test.duration)
		b.AppendBreak(test.strength, test.duration)

		actual := b.Build()
		if actual != test.expected {
			t.Errorf("%s: output mismatch: expected %s, got %s", test.name, test.expected, actual)
		}
	}
}

func TestBuilder_AppendEmphasis(t *testing.T) {
	tests := []struct {
		name     string
		level    emphasis.Level
		text     string
		expected string
	}{
		{
			name:     "default",
			level:    emphasis.Default,
			text:     "test1",
			expected: `<speak><emphasis level="moderate">test1</emphasis><emphasis level="moderate">test1</emphasis></speak>`,
		},
		{
			name:     "strong",
			level:    emphasis.Strong,
			text:     "test2",
			expected: `<speak><emphasis level="strong">test2</emphasis><emphasis level="strong">test2</emphasis></speak>`,
		},
		{
			name:     "moderate",
			level:    emphasis.Moderate,
			text:     "test3",
			expected: `<speak><emphasis level="moderate">test3</emphasis><emphasis level="moderate">test3</emphasis></speak>`,
		},
		{
			name:     "reduced",
			level:    emphasis.Reduced,
			text:     "test4",
			expected: `<speak><emphasis level="reduced">test4</emphasis><emphasis level="reduced">test4</emphasis></speak>`,
		},
		{
			name:     "reduced",
			level:    emphasis.Level("custom"),
			text:     "test5",
			expected: `<speak><emphasis level="custom">test5</emphasis><emphasis level="custom">test5</emphasis></speak>`,
		},
	}

	for _, test := range tests {
		b, _ := NewBuilder()

		b.AppendEmphasis(test.level, test.text)
		b.AppendEmphasis(test.level, test.text)

		actual := b.Build()
		if actual != test.expected {
			t.Errorf("%s: output mismatch: expected %s, got %s", test.name, test.expected, actual)
		}
	}
}

func TestBuilder_AppendParagraph(t *testing.T) {
	b, _ := NewBuilder()

	b.AppendParagraph("text1")
	b.AppendParagraph("text2")

	actual := b.Build()
	expected := `<speak><p>text1</p><p>text2</p></speak>`
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestBuilder_AppendProsody(t *testing.T) {
	tests := []struct {
		name     string
		rate     ProsodyRate
		pitch    ProsodyPitch
		volume   ProsodyVolume
		expected string
	}{
		{
			name:     "x-slow, s-low, & silent",
			rate:     RateXSlow,
			pitch:    PitchXLow,
			volume:   VolumeSilent,
			expected: `<speak><prosody rate="x-slow" pitch="x-low" volume="silent">text1</prosody><prosody rate="x-slow" pitch="x-low" volume="silent">text2</prosody></speak>`,
		},
		{
			name:     "slow, low, & x-soft",
			rate:     RateSlow,
			pitch:    PitchLow,
			volume:   VolumeXSoft,
			expected: `<speak><prosody rate="slow" pitch="low" volume="x-soft">text1</prosody><prosody rate="slow" pitch="low" volume="x-soft">text2</prosody></speak>`,
		},
		{
			name:     "medium, medium, & soft",
			rate:     RateMedium,
			pitch:    PitchMedium,
			volume:   VolumeSoft,
			expected: `<speak><prosody rate="medium" pitch="medium" volume="soft">text1</prosody><prosody rate="medium" pitch="medium" volume="soft">text2</prosody></speak>`,
		},
		{
			name:     "fast, high, & medium",
			rate:     RateFast,
			pitch:    PitchHigh,
			volume:   VolumeMedium,
			expected: `<speak><prosody rate="fast" pitch="high" volume="medium">text1</prosody><prosody rate="fast" pitch="high" volume="medium">text2</prosody></speak>`,
		},
		{
			name:     "x-fast, x-high, & loud",
			rate:     RateXFast,
			pitch:    PitchXHigh,
			volume:   VolumeLoud,
			expected: `<speak><prosody rate="x-fast" pitch="x-high" volume="loud">text1</prosody><prosody rate="x-fast" pitch="x-high" volume="loud">text2</prosody></speak>`,
		},
		{
			name:     "x-fast, x-high, & x-loud",
			rate:     RateXFast,
			pitch:    PitchXHigh,
			volume:   VolumeXLoud,
			expected: `<speak><prosody rate="x-fast" pitch="x-high" volume="x-loud">text1</prosody><prosody rate="x-fast" pitch="x-high" volume="x-loud">text2</prosody></speak>`,
		},
		{
			name:     "custom",
			rate:     ProsodyRate("custom rate"),
			pitch:    ProsodyPitch("custom pitch"),
			volume:   ProsodyVolume("custom volume"),
			expected: `<speak><prosody rate="custom rate" pitch="custom pitch" volume="custom volume">text1</prosody><prosody rate="custom rate" pitch="custom pitch" volume="custom volume">text2</prosody></speak>`,
		},
	}

	for _, test := range tests {
		b, _ := NewBuilder()

		b.AppendProsody(test.rate, test.pitch, test.volume, "text1")
		b.AppendProsody(test.rate, test.pitch, test.volume, "text2")

		actual := b.Build()
		if actual != test.expected {
			t.Errorf("%s: output mismatch: expected %s, got %s", test.name, test.expected, actual)
		}
	}
}

func TestBuilder_AppendSentence(t *testing.T) {
	b, _ := NewBuilder()

	b.AppendSentence("text1")
	b.AppendSentence("text2")

	actual := b.Build()
	expected := `<speak><s>text1</s><s>text2</s></speak>`
	if actual != expected {
		t.Errorf("output mismatch: expected %s, got %s", expected, actual)
	}
}

func TestBuilder_AppendSubstitution(t *testing.T) {
	b, _ := NewBuilder()

	b.AppendSubstitution("alias1", "text1")
	b.AppendSubstitution("alias2", "text2")

	actual := b.Build()
	expected := `<speak><sub alias="alias1">text1</sub><sub alias="alias2">text2</sub></speak>`
	if actual != expected {
		t.Errorf("output mismatch: expected %s, got %s", expected, actual)
	}
}
