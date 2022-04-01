# GameZone - Veb aplikacija bazirana na mikroservisnoj arhitekturi

GameZone je veb aplikacija pomoću koje korisnici imaju mogućnost da pretražuju, filtriraju i kupuju proizvode vezane za gejming kao sto su video igre, konzole, dodatna operema i hardverske komponente i da ostavljaju ocene i komentare za proizvode koje su kupili. Korisnici takođe mogu da prate i komentarisu sve vesti koje su objavljene od strane zaposlenih. Uloge koje postoje u sistemu su registrovani korisnici, zaposleni i administratori.

## Funkcionalnosti aplikacije

### Funkcionalnosti neregistrovanog korisnika

Neregistrovani korisnik može da se registruje i nakon registracije postaje registrovan korisnik sistema. Informacije koje unosi prilikom registracije su korisničko ime, email, ime i prezime. Nakon registracije korisnik prilikom prijave treba da unese verifikacioni kod kojeg dobija na email koji je uneo tokom registracije. Neregistrovani korisnik takođe ima mogućnost da pregleda, pretražuje i fitrira proizvode kao i da pregleda vesti koje postavljaju zaposleni.

### Funkcionalnosti registrovanog korisnika

Registrovani korisnik ima mogućnost da kupuje proizvode koji su dostupni (na lageru). Ako željeni proizvod nije na lageru korisnik se može registrovati da primi email obaveštenje kad taj proizvod postane dostupan. Video igre korisnik može da kupi u fizičkom ili digitalnom obliku. Ukoliko korisnik odluči da kupi igru u digitalnom obliku, na njegov email se šalje šifra kupljene igre. Kada korisnik kupuje proizvod, u okviru glavne stranice nalazi će se lista preporucenih proizvoda koji su slični odabranom proizvodu. On može za sve kupljene proizvode da ostavi ocenu i komentar kako bi i drugi korisnici mogli videti njegovo mišljenje o tim proizvodima. Registrovan korisnik može da pregleda i komentariše vesti koje su objavili zaposleni u okviru aplikacije, a takođe imaju i izbor da se pretplate ili odjave da im vesti stižu na email. Registrovani korisnik može kontaktirati zaposlene ili administratore putem kontakt obrasca. Takođe može i da prijavi druge registrovane korisnike zbog
nedoličnog ponašanja.

### Funkcionalnosti zaposlenih

Zaposleni može da dodaje nove proizvode i da izmeni ili obriše postojeće proizvode. Ima mogućnost da definiše koji proizvodi će biti prikazani u okviru glavne stranice. Takođe, ima mogućnost da piše vesti koje će biti vidljive od strane drugih korisnika u sistemu. Kao i registrovani korisnici i zaposleni može da prijavi registrovane korisnike zbog nedoličnog ponašanja.

### Funkcionalnosti adminstratora

Adminstrator može da dodaje nove zaposlene i administratore u sistem. Može da pregleda sve registrovane korisnike u sistemu i njihovu istoriju kupovine. Takođe može i da formira izveštaje kao što su: broj prodatih igara na osnovu platforme, broj prodatih igara na osnovu oblika ili listu proizvoda koji ostvaruju najveći profit u zadnjih 30 dana. Administrator može da blokira korisnike koji su prijavljeni od strane zaposlenih i drugih registrovanih korisnika na određeni vremenski period.

## Arhitektura sistema

### API Gateway 

Servis koji grupiše sve funkcionalnosti sistema. API Gateway će biti implementiran korišćenjem radnog okvira Flask.

### Mikroservis za korisnike

Mikroservis koji obavlja registraciju, prijavu i dodavanje novih zaposlenih i administratora. Korisnički mikroservis će biti implementiran korišćenjem radnog okvira Flask.

### Mikroservis za slanje email-a

Mikroservis koji pruža slanje email-ova. Mikroservis za slanje email-a će biti implementiran korišćenjem radnog okvira Flask.

### Mikroservis za proizvode

Ovaj mikroservis omogućava pretraživanje, filtriranje, dodavanje, izmenu i brisanje proizvoda i definisanje koji proizvodi će se nalaziti u okviru glavne stranice. Takođe, ovaj mikroservis obavlja sve funkcionalnosti u vezi kupovine proizvoda. Mikroservis za proizvode će biti implementiran korisćenjem jezika Go.

### Mikroservis za ocenjivanje i komentarisanje kupljenih proizvoda

Ovaj mikroservis pruža registrovanim korisnicima mogućnost ocenjivanja i komentarisanja kupljenih proizvoda. Mikroservis za ocenjivanje i komentarisanje kupljenih proizvoda će biti implementiran korišćenjem jezika Go.

### Mikroservis za vesti

Ovaj mikroservis pruža mogućnost da se registrovani korisnici pretplate ili odjave od dobijanja vesti na email, kao i da zaposleni postavljaju vesti u okviru aplikacije, koje nakon toga registrovani korisnici mogu pregledati i komentarisati. Mikroservis za vesti će biti implementiran korišćenjem jezika Go.

### Mikroservis za kontakt i prijavu nedoličnog ponašanja

Ovaj mikroservis pruža mogućnost da registrovani korisnici i zaposleni prijave druge registrovane korisnike zbog nedoličnog ponašanja. Takođe ovaj mikroservis obavlja sve funkcionalnosti vezane za kontaktiranje zaposlenog ili administratora putem kontakt obrasca. Mikroservis za kontakt i prijavu nedoličnog ponašanja će biti implementiran korišćenjem jezika Go.

### Mikroservis za izveštaje

Ovaj mikroservis pruža formiranje izveštaja od strane administratora. Mikroservis za izvestaje će biti implementiran korišćenjem okruženja Pharo.

### Klijentska aplikacija

Klijentska aplikacija omogućava pristup svim prethodno navedenim funkcionalnostima. Klijentska aplikacija će biti implementirana korišćenjem React JavaScript biblioteke.
