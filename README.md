# GameZone - Veb aplikacija bazirana na mikroservisnoj arhitekturi

GameZone je veb aplikacija pomocu koje korisnici imaju mogucnost da pretrazuju, filtriraju i kupuju proizvode vezane za gejming kao sto su video igre, konzole, dodatna operema i hardverske komponente i da ostavljaju ocene i komentare za proizvode koje su kupili. Korisnici takodje mogu da prate i komentarisu sve vesti koje su objavljene od strane zaposlenih. Uloge koje postoje u sistemu su registrovani korisnici, zaposleni i administratori.

## Funkcionalnosti aplikacije

### Funkcionalnosti neregistrovanog korisnika

Neregistrovani korisnik moze da se registruje i nakon registracije postaje registrovan korisnik sistema. Informacije koje unosi prilikom registracije su korisnicko ime, email, ime i prezime. Nakon registracije korisnik prilikom prijave treba da unese verifikacioni kod kojeg dobija na email koji je uneo tokom registracije. Neregistrovani korisnik takodje ima mogucnost da pregleda, pretrazuje i fitrira proizvode kao i da pregleda vesti koje postavljaju zaposleni.

### Funkcionalnosti registrovanog korisnika

Registrovani korisnik ima mogucnost da kupuje proizvode koji su dostupni (na lageru). Ako zeljeni proizvod nije na lageru korisnik se moze registrovati da primi email obavestenje kad taj proizvod postane dostupan. Video igre korisnik moze kupi u fizickom ili digitalnom obliku. Ukoliko korisnik odluci da kupi igru u digitalnom obliku, na njegov email se salje sifra kupljene igre. Kada korisnik kupuje proizvod, u okviru glavne stranice nalazi ce se lista preporucenih proizvoda koji su slicni odabranom proizvodu. On moze za sve kupljene proizvode da ostavi ocenu i komentar kako bi i drugi korisnici mogli videti njegovo misljenje o tim proizvodima. Registrovan korisnik moze da pregleda i komentarise vesti koje su objavili zaposleni u okviru aplikacije, a takdoje imaju i izbor da se pretplate ili odjave da im vesti stizu na email. Registrovani korisnik moze kontaktirati zaposlene ili administratore putem kontakt obrasca.

### Funkcionalnosti zaposlenih

Zaposleni moze da dodaje nove proizvode i da izmeni ili obrise postojece proizvode. Ima mogucnost da definise koji proizvodi ce biti prikazani u okviru glavne stranice. Takodje, ima mogucnost da pise vesti koje ce biti vidljive od strane drugih korisnika u sistemu.

### Funkcionalnosti adminstratora

Adminstrator moze da dodaje nove zaposlene i administratore u sistem. Moze da pregleda sve registrovane korisnike u sistemu i njihovu istoriju kupovine. Takodje moze i da formira izvestaje kao sto su: broj prodatih igara na osnovu platforme, broj prodatih igara na osnovu oblika ili listu proizvoda koji ostvaruju najveci profit u zadnjih 30 dana. 

## Arhitektura sistema

### API Gateway 

Servis koji grupise sve funkcionalnosti sistema. API Gateway ce biti implementiran koriscenjem radnog okvira Flask.

### Mikroservis za korisnike

Mikroservis koji obavlja registraciju, prijavu i dodavanje novih radnika i administratora. Korisnicki mikroservis ce biti implementiran koriscenjem radnog okvira Flask.

### Mikroservis za slanje email-a

Mikroservis koji pruza slanje email-ova. Mikroservis za slanje email-a ce biti implementiran koriscenjem radnog okvira Flask.

### Mikroservis za proizvode

Ovaj mikroservis omogucava pretrazivanje, filtriranje, dodavanje, izmenu i brisanje proizvoda i definisanje koji proizvodi ce se nalaziti u okviru glavne stranice. Takodje, ovaj mikroservis obavlja sve funkcionalnosti u vezi kupovine proizvoda. Mikroservis za proizvode ce biti implementiran koriscenjem jezika Go.

### Mikroservis za ocenjivanje i komentarisanje

Ovaj mikroservis pruza registrovanim korisnicima mogucnost ocenjivanja i komentarisanja kupljenih proizvoda. Mikroservis za ocenjivanje i komentarisanje ce biti implementiran koriscenjem jezika Go.

### Mikroservis za vesti

Ovaj mikroservis pruza mogucnost da se registrovani korisnici pretplate ili odjave od dobijanja vesti na email, kao i da radnici postavljaju vesti u okviru aplikacije, koje nakon toga registrovani korisnici mogu pregledati i komentarisati. Mikroservis za vesti ce biti implementiran koriscenjem jezika Go.

### Mikroservis za izvestaje

Ovaj mikroservis pruza formiranje izvestaja od strane administratora. Mikroservis za izvestaje ce biti implementiran koriscenjem okruzenja Pharo.

### Klijentska aplikacija

Klijentska aplikacija omogucava pristup svim prethodno navedenim funkcionalnostima. Klijentska aplikacija ce biti implementirana koriscenjem React JavaScript biblioteke.
