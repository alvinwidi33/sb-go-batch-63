## ROLE ADMIN
REGISTER AND LOGIN
{
    "username":"ayam",
    "password":"ayam"
}

POST SALOON
{
  "id":1,
  "name": "Luxury Saloon90",
  "location": "Bogor",
  "open": "2025-01-31T09:00:00Z",
  "close": "2025-01-31T21:00:00Z"
}

PUT SALOON
{
  "name": "Luxury Saloon",
  "location": "Bogor",
  "open": "2025-01-31T09:00:00Z",
  "close": "2025-01-31T21:00:00Z"
}

POST RESERVATION
{ 
  "id":0,
  "services": {
    "0": "meni",
    "1": "pedi"
  },
  "start": "2025-01-31T10:00:00Z",
  "customer_id": "f05e1dfa-70d2-4e3e-ac38-08d431e7cd15",
  "saloon_id": 1
}

DONE RESERVATION
{
    "rating":4,
    "feedback":"bagus"
}