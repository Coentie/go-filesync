**Entiteiten definities**
**Een taak** is een actie die uitgevoerd moet worden. 
- Een taak moet een naam hebben om te communiceren welke actie er ondernomen moet worden
- Een taak erkent de volgende statussen
	- Idle (moet opgepakt worden)
	- Busy (is opgepakt)
	- Done (afgerond)
- Een taak kan een gebruiker hebben
- Een taak MOET een type hebben
	- Boolean (aanvinken dat iets gebeurd is)
	- Mail versturen (kan automatisch)
	- Formulier invullen

**Een fase** is een verzameling aan taken en een segment uit een **takenlijst**. 
- Een fase KAN een afhankelijkheid hebben naar een andere fase.
- De afhankelijkheid van een fase moet instelbaar zijn op taakniveau
- By default heeft een fase geen afhankelijkheid aan elkaar om zo de vrijheid van gebruik te garanderen. 

**Een takenlijst** is een verzameling aan fases. De takenlijst kan aan verschillende entiteiten gehangen worden die het concept van "beginnen en afronden" (vanaf nu een Startable genoemd) kennen. 

**Keuzes**
- 1 of meerdere takenlijsten aan een entiteit kunnen hangen. 
	- Voordeel van meerdere:
		- Je kunt compositie toepassen op je takenlijsten waar je een aantal kleinere lijsten maakt en die op maat hangt aan Startable waar zij op toepassen. 
		- Je hoeft niet, wanneer je afwijkt van je normale takenlijst, een hele nieuwe te maken. 
	- Nadeel:
		- Je kunt niet automatisch bepalen of een Startable afgerond is.
		- Het kan onoverzichtelijk worden als je veel verschillende takenlijst hebt