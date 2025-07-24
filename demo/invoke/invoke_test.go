package invoke

import (
	"fmt"
	"testing"
)

func TestHttpAiInvoke(t *testing.T) {
	fmt.Println("=====================HttpAiInvoke=====================")
	http_ai_invoke()
}

func TestEinoAiInvoke(t *testing.T) {
	fmt.Println("=====================EinoAiInvoke=====================")
	eion_ai_invoke()
}

func TestSdkAiInvoke(t *testing.T) {
	fmt.Println("=====================SdkAiInvoke=====================")
	sdk_ai_invoke()
}

func TestLangChainGoAiInvoke(t *testing.T) {
	fmt.Println("====================LangChainGoAiInvoke====================")
	langchaingo_ai_invoke()
}
