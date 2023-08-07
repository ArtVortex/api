package contentStore

import (
	"artvortex-api/internal/config"
	"bytes"
	"io/ioutil"

	shell "github.com/ipfs/go-ipfs-api"
)

func IPFSGetByCID(cid string, ipfsNodeAddress string)([]byte, error) {
	nftConfig := config.NFTServiceConfig

	if (ipfsNodeAddress == "") {
		ipfsNodeAddress = nftConfig.IPFSHost + ":" + nftConfig.IPFSPort
	}

	ipfsShell := shell.NewShell(ipfsNodeAddress)

	// Get IPFS object reader
	reader, err := ipfsShell.Cat(cid)

	if err != nil {
		return nil, err
	}
	//close ready after function completion
	defer reader.Close()

	//Read content of IPFS object
	content, err := ioutil.ReadAll(reader)

	if err != nil {
			return nil, err
	}

	return content, nil
}

/**
* - adds content to ipfs node using Add method. 
* - Add is more suitable for simple object while Dag is more flexible
* for complex content which has links to other ipfs object. eg. array of ipfs objects
*/ 
func IPFSAddContent(content []byte, ipfsNodeAddress string)(string, error) {
	nftConfig := config.NFTServiceConfig

	if ipfsNodeAddress == "" {
		ipfsNodeAddress = nftConfig.IPFSHost + ":" + nftConfig.IPFSPort
	}
	ipfsShell := shell.NewShell(ipfsNodeAddress)
	cid, err := ipfsShell.Add(bytes.NewReader(content))
	if err != nil {
		return "", err
	}

	return cid, nil
}

func IPFSAddContentFromURL(url string, ipfsNodeAddress string) (string, error) {
	// Get content from URL
	content, err := GetObjectFromURL(url)
	if err != nil {
		return "", err
	}

	// Add content to IPFS node
	cid, err := IPFSAddContent(content, ipfsNodeAddress)
	if err != nil {
		return "", err
	}

	return cid, nil
}