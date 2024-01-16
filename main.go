package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/minias/sss-go/sss"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: sss-go [64 bytes of private key]\n")
	fmt.Fprintf(os.Stderr, "example:\n")
	fmt.Fprintf(os.Stderr, "./sss-go 0123456789abcdefghijklmnopqrstuvwxyz0123456789abcdefghijklmnopqr\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}
	data := []byte(args[0])
	//data := []byte("0123456789abcdefghijklmnopqrstuvwxyz0123456789abcdefghijklmnopqr")

	log.Printf("private key: \n%s\n", data)
	//log.Printf("data : %x\n", data)

	log.Printf("shares : \n%s\n", "CreateShares 2 of 3")
	// Create 5 shares; allow 4 to restore the original data
	//shares, err := sss.CreateShares(data, 5, 4)
	shares, err := sss.CreateShares(data, 3, 2) // 2 of 3
	if err != nil {
		log.Fatalln(err)
	}
	//log.Printf("shares : %s\n", shares)
	log.Printf("CreateShares[0] : \n%x\n", shares[0])
	//012850a4e198d6f9c94b05a63439c00bfda864f90c5f5264674958d814a8c5a7c29285a9cbe7a01283abe90da020358d6f8f5e870b3b02d500ca11ac051f351820949c7622b413243dc58e7a7bde1735a4f9fb8903dd1949e5023fd40de7ce6e2f186d59a590bbcde5b564d0249707050a
	log.Printf("CreateShares[1] : \n%x\n", shares[1])
	//021ad43b3b97057336b856b8bba149fc4f5bea086d11947f8fd73238975e9bce519285a9cbe7a01283abe90da020358d6f8f5e870b3b02d500ca11ac051f351820949c7622b413243dc58e7a7bde1735a4f9fb8903dd1949e5023fd40de7ce6e2f186d59a590bbcde5b564d0249707050a
	log.Printf("CreateShares[2] : \n%x\n", shares[2])
	//03fda84e8492bdfc63e967b23720c758210a90aebb2bd676d75414911f0c58e9209285a9cbe7a01283abe90da020358d6f8f5e870b3b02d500ca11ac051f351820949c7622b413243dc58e7a7bde1735a4f9fb8903dd1949e5023fd40de7ce6e2f186d59a590bbcde5b564d0249707050a
	//log.Println(shares)

	// Permute and lose some of the shares (for demonstrational purposes)
	// new_shares := make([][]byte, 4)
	// new_shares[0] = shares[2]
	// new_shares[1] = shares[4]
	// new_shares[2] = shares[0]
	// new_shares[3] = shares[3]

	// 2of3
	new_shares := make([][]byte, 2)
	// case 1
	// new_shares[0] = shares[1]
	// 023c60c4d3d09631320d0a381383bfc527bcd089807ca096d69002fa138fa91b2bdceb9c26f4287829e734dff7cae93ed1b521bf0484263ca8302c24ffea1fd30022e5ae03863a893930218679c43a1ae1f1ea31070d119e2d16e03ac249a73a267f0211f283ed1d25194259d224491562
	// new_shares[1] = shares[2]
	// 0325dd12ce79969dca79816af1472b6d6e4ba27f03a51adbc8cbd00c8f13db21d6dceb9c26f4287829e734dff7cae93ed1b521bf0484263ca8302c24ffea1fd30022e5ae03863a893930218679c43a1ae1f1ea31070d119e2d16e03ac249a73a267f0211f283ed1d25194259d224491562

	// case 2
	// new_shares[0] = shares[0]
	// 011bbc82468a869acaaf321cb5637e9f64c11fd4360e7cbe9852e8c7d5f907c7bc539498beaac77ad64973f897e7953ae7d93dfad4700fef237d62748c3fabd34cffbef78d88880cc16c988a58dd0857813b21c46c8e5477fb1e9a26142fee95fa09a3837e57e3f655e05e33618589b687
	// new_shares[1] = shares[2]
	// 03cdbabb8dda46c08821b9bc1ab0d93cecd5dff107cec5165625bce84bc8025a9f539498beaac77ad64973f897e7953ae7d93dfad4700fef237d62748c3fabd34cffbef78d88880cc16c988a58dd0857813b21c46c8e5477fb1e9a26142fee95fa09a3837e57e3f655e05e33618589b687

	// case 3
	new_shares[0] = shares[1]
	// 021ad43b3b97057336b856b8bba149fc4f5bea086d11947f8fd73238975e9bce519285a9cbe7a01283abe90da020358d6f8f5e870b3b02d500ca11ac051f351820949c7622b413243dc58e7a7bde1735a4f9fb8903dd1949e5023fd40de7ce6e2f186d59a590bbcde5b564d0249707050a
	new_shares[1] = shares[0]
	// 012850a4e198d6f9c94b05a63439c00bfda864f90c5f5264674958d814a8c5a7c29285a9cbe7a01283abe90da020358d6f8f5e870b3b02d500ca11ac051f351820949c7622b413243dc58e7a7bde1735a4f9fb8903dd1949e5023fd40de7ce6e2f186d59a590bbcde5b564d0249707050a

	//log.Printf("new_shares : %s\n", new_shares)
	log.Printf("combine shares[0] : \n%x\n", new_shares[0])
	//021ad43b3b97057336b856b8bba149fc4f5bea086d11947f8fd73238975e9bce519285a9cbe7a01283abe90da020358d6f8f5e870b3b02d500ca11ac051f351820949c7622b413243dc58e7a7bde1735a4f9fb8903dd1949e5023fd40de7ce6e2f186d59a590bbcde5b564d0249707050a
	log.Printf("combine shares[1] : \n%x\n", new_shares[1])
	//012850a4e198d6f9c94b05a63439c00bfda864f90c5f5264674958d814a8c5a7c29285a9cbe7a01283abe90da020358d6f8f5e870b3b02d500ca11ac051f351820949c7622b413243dc58e7a7bde1735a4f9fb8903dd1949e5023fd40de7ce6e2f186d59a590bbcde5b564d0249707050a
	//log.Println(new_shares)

	// Try to restore the original secret
	restored, err := sss.CombineShares(new_shares)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("restored key: \n%s\n", restored)
	//0123456789abcdefghijklmnopqrstuvwxyz0123456789abcdefghijklmnopqr
	//log.Printf("restored : %x\n", restored)
}
