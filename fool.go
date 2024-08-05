package main

import "os"
import "bufio"
import "fmt"
import "sort"
import "strconv"
import "strings"
import "reflect"

import "math/rand"

//import "rsc.io/quote"

type Card struct {
    sign string
    val  string
    weight int
}
type Player struct {
    name string
    cards []Card
    human bool
}

type Game struct {
    players []Player
    current int
    next int
    kozer Card
    table []Card
    byta []Card
}

var pcount = 4;
var dec = [13]string{"2","3","4","5","6","7","8","9","10","J","Q","K","A"};
var sign = [4]string{"♡","♧","♤","♢"};
var signtext = [4]string{"čirvų","kryžių","vynų","būgnų"};
var pnames = []string{"jonas","petras","kazys","albinas"};

//var gamedec = [52]string{};
var gdec = [52]Card{};	
var gcur = 0;
var g Game
var pn string

func main() {
    //fmt.Println(quote.Go())
    fmt.Println("Hello, fool durak World!")
    fmt.Println("press enter to skip or enter name")
    name := uin("name: ")
    if len(name) > 0 {
	pn = name
	addplayer(name)
    }

    initcards()
    shuffle()
    startgame()
    //win := 0
    //for gcur < 52 {
  //for win == 0 {
    empty := 0
  for empty < pcount -1{
    if g.current == 0 {
	    //empty = 0
    }

    fmt.Println("gcur = ",gcur)
    fmt.Println("kozer = ",g.kozer.sign, g.kozer.val)
    fmt.Println("current",g.current)

    //fmt.Println("game stuff")
    empty = 0
    for i, p := range g.players {
      fmt.Println(i,p.name,len(p.cards))
      if len(p.cards) == 0 { empty++ }
    }



    if len(g.players[g.current].cards) == 0 {
	
      //fmt.Println("empty++")
	//win = g.current
	//empty++
        g.current = getnext(g.current)
	continue
	//break
    }
    pnext := getnext(g.current)
    //if len(g.players[pnext].cards) == 0 {
    for len(g.players[pnext].cards) == 0 {
	//win = pnext
	//empty++

        pnext = getnext(pnext)
	//continue
	//break
    }
    if g.current == pnext {
	    break
    }
    g.next = pnext

    g = play(g, 0)
    
    for i := 0; i < pcount; i++ {
        g.players[i]=deal(g.players[i])
    }

    fmt.Println("precurrent", g.current)
    g.current = getnext(g.current)
    fmt.Println("postcurrent", g.current)

    fmt.Println("empty", empty)
  }//if <52

    fmt.Println("end empty", empty)

}

func getnext(current int) int {
    pnext := current+1
    if pnext > pcount-1 { pnext = pnext-pcount }
    return pnext
}

func addplayer(name string) {
    pcount++
    pnames = append(pnames, name)
    
}

func uin(out string) string{
    fmt.Print(out)
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)
    fmt.Println("received: ",text)
    return text
}

func initcards(){
    for ii, vv := range sign {   
        for i, v := range dec {   
            //fmt.Println(i, v, "of",vv, ii)
	    card := Card{vv,v,i}
	    gdec[i+13*ii] = card;
	    //gamedec[i+13*ii] = v+vv;
        } 
    }
}

func shuffle(){
    rand.Shuffle(len(gdec), func(i, j int) {
        gdec[i], gdec[j] = gdec[j], gdec[i]
    })
    //fmt.Println(gdec)
}

func nextcard() Card{

    fmt.Println(gcur, "gcur nc")
    if gcur == 52 {
	//fmt.Println(52)
        gcur++;
        return g.kozer 
    }
    if gcur >= 53 {
        return Card{} 
    }
    next := gdec[gcur]
    gcur++;
    if g.kozer.sign == next.sign {
	next.weight = next.weight +13
        fmt.Println("next",next)
    }
    return next
}

func sortic(a []Card) {
    sort.Slice(a, func(i, j int) bool {
        return a[i].weight < a[j].weight
    })
}

func deal(p Player) Player{ 
    for i := len(p.cards); i < 6; i++ {
        if gcur >= 53 {
	    return p
        }
	next := nextcard()
	p.cards = append(p.cards, next)
        //fmt.Println(p.cards)
    }
    sortic(p.cards)
    return p
}

func startgame() Game{
    g = Game{}
    g.kozer = nextcard()
    g.kozer.weight += 13
    for i := 0; i < pcount; i++ {
	p := Player{};
	p.name = pnames[i];
	if p.name == pn {
		p.human = true
	}else{
		p.human = false
	}
	p = deal(p);
        //fmt.Println(p)
	g.players = append(g.players, p);
        //fmt.Println(g.players)
    }
    smallest := Card{}
    smallest.weight = 19;


    fmt.Println("pcount",pcount)
    //who starts
    for i := 0; i < pcount; i++ {
	k1 := findkozer(g.players[i].cards, g.kozer.sign)
        fmt.Println(k1)
	if smallest.weight > k1.weight {
	    smallest = k1
	    g.current = i
        }
    }

    return g
}

func findkozer(c []Card, kozer string) Card {
    ind := findkozerind(c, kozer);
    if ind < 0 {
        return Card{kozer, "20", 40}
    }
    return c[ind]
}

func findkozerind(c []Card, kozer string) int {
    for i := 0; i < len(c); i++ {
	if c[i].sign == kozer {
	    return i
        }
    }
    return -1
}

func findhigher(cs []Card, c Card) int {
    for i := 0; i < len(cs); i++ {
	if cs[i].sign == c.sign && cs[i].weight > c.weight {
	    return i
        }
    }
    return -1
}

func remove(s []Card, i int) []Card {
    s[i] = s[len(s)-1]
    // We do not need to put s[i] at the end, as it will be discarded anyway
    return s[:len(s)-1]
}


func findsameval(c []Card, cf []Card) []int {
    same := []int{}
    for i := 0; i < len(c); i++ {
        for ii := 0; ii < len(cf); ii++ {
	    if c[i].val == cf[ii].val {
		same = append(same, i)
	    }

	}  
    }
    return same
}

func princards(c []Card){
    fmt.Println("===============")
    for ind, card := range c {
	fmt.Println("ind", ind, "card", card.sign, card.val)
    
    }	
}

func play(g Game, cardind int) Game{

    //chech first card non korez
    //if kozer go next,
    //but if all cards korers go with 0


    //fmt.Println("its here", cardind, "g",g)
    //take fitst, smallest from sorted 
    if g.players[g.current].human == true {
	    princards(g.players[g.current].cards)
	    tin := uin("enter index of card to add: ")
	    tin = strings.TrimSpace(tin)
	    cardindi, _ := strconv.Atoi(tin)
	    if tin != "" {
	        cardind = cardindi
            }
	    fmt.Println("cardind",cardind)
    }
    if cardind < 0 {
	fmt.Println(" byta ")
	return g
    }
    g.table = append(g.table, 
        g.players[g.current].cards[cardind])
    g.players[g.current].cards = 
        remove(g.players[g.current].cards, cardind)
    sortic(g.players[g.current].cards)
    fmt.Println(g.current,"g start play table")
    princards( g.table)

    //pnext := getnext(g.current)
    pnext := g.next
    //index of card to respond higher val 
    //and same sign
    ind := 
        findhigher(g.players[pnext].cards, 
	g.table[len(g.table)-1]    )

    if ind < 0 && g.table[len(g.table)-1].sign != g.kozer.sign{ //not found higher card
        //find kozer
        ind = findkozerind(g.players[pnext].cards, 	    g.kozer.sign)
    }
    fmt.Println(ind, "ind of response card found")
    if g.players[pnext].human == true {
	    princards(g.players[pnext].cards)
	    tin := uin("enter index of card to respond: ")
	    fmt.Println(" tin= ", tin)
	    fmt.Println(reflect.TypeOf(tin))
	    tin = strings.TrimSpace(tin)
	    if tin != "" {
	        cardindi, err := strconv.Atoi(tin)
	        ind = cardindi
	        fmt.Println(" tin err= ", err)
	        fmt.Println(" cardindi= ", cardindi)
            }
	    fmt.Println(" ind= ", ind)
    }
    if ind < 0 { //not found higher card
	//go home
	fmt.Println("go home");
	g.players[pnext].cards = append(g.players[pnext].cards, g.table...);
        sortic(g.players[pnext].cards)
        g.table = []Card{}
	fmt.Println("go home pnext",pnext);
	//pnext = pnext+1
        //if pnext > pcount-1 { pnext = pnext-pcount }
	g.current = pnext
	fmt.Println("go home cur ",g.current);
	return g
    }else{
        g.table = append(g.table, 
            g.players[pnext].cards[ind])
        g.players[pnext].cards = 
            remove(g.players[pnext].cards, ind)
        sortic(g.players[pnext].cards)
        fmt.Println("g resp play table")
	princards( g.table)
    }


    //if has to add same val cards
    same := 
        findsameval(g.players[g.current].cards,
        g.table )
    fmt.Println(same, " same ind ")
    //for i := 0; i < len(same); i++ {

	//recursive play with id of same 
	//instead of 0
    //}

    if len(same) < 1 {
        g.table = []Card{}
        fmt.Println( " byta")
    }else{

        fmt.Println(same[0], " same dametimas ")
	if len( g.players[g.current].cards) > 0 {
	if len( g.players[pnext].cards) > 0 {
           g = play(g, same[0])
        }
        }

    }
    return g
}
