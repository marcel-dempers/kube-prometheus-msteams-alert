package models

type Alerts struct {
	Alerts []Alert
}

type Alert struct {
   Status string
   Annotations Annotations
   Labels Labels
}

type Annotations struct {
	Description string
	Summary string
}

type Labels struct {
	Cluster string
}