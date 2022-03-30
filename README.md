# GameZone - Veb aplikacija bazirana na mikroservisnoj arhitekturi

GameZone je veb aplikacija na kojoj korisnici imaju mogucnost da pretrazuju, filtriraju i kupuju proizvode vezane za gejming kao sto su video igre, konzole, dodatna operema i hardverske komponente i da ostavljaju kritike za proizovode koje su kupili. Korisnici takodje mogu da prate i kometarisu na sve vesti koje su objavljene od strane zaposlenih. Uloge koje postoje u sistemu su registrovani korisnici, radnici i administratori.

## Funkcionalnosti aplikacije

### Funkcionalnosti neregistrovanog korisnika

Neregistrovani korisnik moze da se registruje i prijavi na sistem. Informacije koje unosi prilikom registracije su korisnicko ime, email, ime i prezime. Nakon registracije korisnik prilikom prijave treba da unese verifikacioni kod koji dobije na email koji je uneo tokom registracije. Neregistrovani korisnik takodje ima mogucnost da pregleda, pretrazuje i fitrira proizvode kao i da pregleda vesti koje postavljaju radnici.

### Funkcionalnosti registrovanog korisnika

Registrovani korisnik ima mogucnost da kupuje proizvode koji su dostupni. Video igre korisnik moze kupi u fizickom ili digitalnom obliku. Ukoliko korisnik odluci da kupi igru u digitalnom obliku, na njegov email se salje sifra kupljene igre. Korisnik kada kupi proizvod, u okviru glavne stranice nalazi ce se lista preporucenih proizvoda koji su slicni onim prozivodima koje je ranije kupovao. Moze za sve kupljene proizvode da ostavi ocenu i komentar kako bi mogli drugi korisnici mogli videti njegovo miseljenje o tim proizvodima. Registrovan korisnik moze da pregleda i komentarise na vestima koji su objavili radnici u okviru aplikacije, a takdoje imaju i izbor da se pretplate/odjave da im stizu vesti na email.

### Funkcionalnosti radnika

Radnik moze da dodaje nove proizvode i da izmeni ili obrise postojece proizvode. Ima mogucnost da definise koji proizvodi ce biti prikazani u okviru glavne stranice. Imaju mogucnost da pisu vesti koji cu biti vidljivi od strane drugih korisnika u sistemu.

### Funkcionalnosti adminstratora

Adminstrator moze da dodaje nove radnike i administratore u sistem. Moze da pregleda sve registrovane korisnike u sistemu i njihovu istoriju kupovine. Takodje moze i da formira izvestaje kao sto su: broj prodatih igara na osnovu platforme, broj prodatih igara na osnovu oblika i proizovodi koji ostvaruju najveci profit. 

## Arhitektura sistema

### API Gateway 

Servis koji grupise sve funkcionalnosti sistema. API Gateway ce biti implementiran koriscenjem Flask.

### Mikroservis za korisnike

Mikroservis koji obavlja registraciju, prijavu i dodavanje novih radnika i administratora. Korisnicki mikroservis ce biti implementiran koriscenjem Flask.

### Mikroservis za slanje email-a

Mikroservis koji pruza slanje email-ova. Mikroservis za slanje email-a ce biti implementiran koriscenjem Flask.

### Mikroservis za proizvode

Mikroservis koji omogucava pretrazivanje, filtriranje, dodavanje, izmena i brisanje proizvoda i definisanje koji proizvodi ce se nalaziti u okviru glavne stranice. Mikroservis za proizvode ce biti implementiran koriscenjem Go.

### Mikroservis za ocenivanje i komentarisanje

Mikroservis koji pruza registrovanim korisnicima mogucnost ocenjivanja i kometarisanja na kupljene proizvode. Mikroservis za ocenivanje i komentarisanje ce biti implementiran koriscenjem Go.

### Mikroservis za vesti

Mikroservis koji pruza mogucnost da registrovani korisnici se pretplate/odjave od dobijenja vesti na email i radnicima da postavljaju vesti u okviru aplikacije koji registrovani korisnici mogu da pregledaju i da komentarisu na njima. Mikroservis za vesti ce biti implementiran koriscenjem Go.

### Mikroservis za pregled korisnika, istorije kupovine i izvestaje

Mikroservis koji pruza formiranje izvestaja i pregled svih registrovanih korisnika kao i njihovu istoriju kupovina. Mikroservis za pregled korisnika, istorije kupovine i izvestaje ce biti implementiran koriscenjem Pharo.

### Klijentska aplikacija

Klijentska aplikacija omogucava nam pristup svim prethodno navedenim funkcionalnostima. Kljentska aplikacija ce biti implementirana koriscenjem React.
