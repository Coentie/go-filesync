**Doel:** Het doel van de taskflow is het aanbieden van een voorafgedefineerde lijst met actiepunten en op basis van afgeronde en niet-afgeronde actiepunten een status te bepalen. 

**Eisen**
- De taskflow moet op een eenvoudige manier de huidige status van een entiteit kunnen bepalen.
- De taskflow moet de mogelijkheid hebben taken te kunnen openstellen / blokkeren op basis van de status van andere taken.
- De taskflow moet op chronologische wijze worden afgerond. Dat betekend niet dat de afzonderlijke taken ook op chronologische wijze moeten worden afgerond. 

**Entiteit:**  Taskflow
**Verantwoordelijkheden:**
- De huidige taak serveren
	- Waarin de huidige taak de eerste taak op chronologische volgorde is of geen taak.
	- De huidig openstaande taak = de status, of als er geen openstaande taak meer is, dan is de status afgerond. 

**Entiteit:** Taak
**Verantwoordelijkheden**
- Beschrijven aan de gebruiker wat gedaan moet worden, of wat al is gedaan
**Eigenschappen**
- Naam
- 1 of meerdere gebruikers
- Type (in de toekomst omtrent automatisering)
	- E-mail
	- Actie
	- Formulier invullen
- 4 Statussen
	- Idle (klaar om te beginnen)
	- Busy (Wordt aan gewerkt)
	- Done (afgerond)
	- Blocked (Kan pas aan worden begonnen als een andere taak is afgerond)
- 