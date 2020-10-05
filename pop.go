package main

import (
	"math/rand"
	"time"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/services"
	"github.com/Moletastic/utem-gsp/store"
	"github.com/jinzhu/gorm"
)

func pop(d *gorm.DB) error {
	ac := store.NewAccessStore(d)
	es := store.NewEducationStore(d)
	pro := store.NewProjectStore(d)
	department := models.Department{
		Name: "Departamento de Informática y Computación",
	}
	es.Related[1].Service.Create(&department)
	careers, _ := generateCareers(es.Related[0].Service, department)
	students, err := generateStudents(es.Related[2].Service, careers)
	if err != nil {
		return err
	}
	teachers, err := generateTeachers(ac.Related[0].Service)
	if err != nil {
		return err
	}
	subjects, err := generateSubjects(pro.Related[4].Service)
	if err != nil {
		return err
	}
	states, err := generateProjectStates(pro.Related[12].Service)
	if err != nil {
		return err
	}
	types, err := generateProjectTypes(pro.Related[11].Service)
	if err != nil {
		return err
	}
	ltypes, err := generateLinkTypes(pro.Related[0].Service)
	if err != nil {
		return err
	}
	chs, err := generateChannels(pro.Related[8].Service)
	projects, err := generateProjects(pro.Project, students, teachers, subjects, types, states)

	if err != nil {
		return err
	}

	rubrics, err := generateRubric(pro.Related[9].Service)

	if err != nil {
		return err
	}

	for _, project := range projects {
		err = fillProject(pro.Project, &project, rubrics, chs, ltypes)
		if err != nil {
			return err
		}
	}

	_, err = generateAdmins(ac.Related[1].Service)

	if err != nil {
		return err
	}

	return nil
}

type FLName struct {
	F string
	L string
}

func generateCareers(s *services.CRUDService, department models.Department) ([]models.Career, error) {
	cs := []models.Career{
		{
			Code:         21041,
			DepartmentID: department.ID,
			Name:         "Ingeniería en Computación Mención Informática",
		},
		{
			Code:         21030,
			DepartmentID: department.ID,
			Name:         "Ingeniería en Informática",
		},
		{
			Code:         21049,
			DepartmentID: department.ID,
			Name:         "Ingeniería en Ciencias de Datos",
		},
	}
	careers := make([]models.Career, 0)
	for _, c := range cs {
		err := s.Create(&c)
		if err != nil {
			return nil, err
		}
		careers = append(careers, c)
	}
	return careers, nil
}

func generateStudents(s *services.CRUDService, careers []models.Career) ([]models.Student, error) {

	sts := []models.Student{
		models.NewStudent("Gabriel", "Marinan", careers[0].ID),

		models.NewStudent("Alejandra", "Munoz", careers[0].ID),
		models.NewStudent("Michel", "Hernandez", careers[0].ID),
		models.NewStudent("Camilo", "Robles", careers[0].ID),

		models.NewStudent("Jordan", "Jimenez", careers[0].ID),
		models.NewStudent("Nicolas", "Martinez", careers[0].ID),
		models.NewStudent("Marcelo", "Diaz", careers[0].ID),

		models.NewStudent("Kevin", "Alarcon", careers[0].ID),
		models.NewStudent("Sergio", "Licanqueo", careers[0].ID),
		models.NewStudent("Esteban", "Martinez", careers[0].ID),

		models.NewStudent("Ignacio", "Valdes", careers[0].ID),
		models.NewStudent("Nicolas", "Alarcon", careers[0].ID),
		models.NewStudent("Rodrigo", "Lobos", careers[0].ID),

		models.NewStudent("Benjamin", "Vargas", careers[0].ID),
		models.NewStudent("Gina", "Garrido", careers[0].ID),
		models.NewStudent("Omar", "Gutierrez", careers[0].ID),

		models.NewStudent("Juan", "Villalobos", careers[0].ID),
		models.NewStudent("Boris", "Vasquez", careers[0].ID),
		models.NewStudent("Alexandra", "Olivares", careers[0].ID),

		models.NewStudent("Alejandro", "Yanez", careers[0].ID),
		models.NewStudent("Cristobal", "Morales", careers[0].ID),
		models.NewStudent("Dixon", "Ortiz", careers[0].ID),

		models.NewStudent("Renato", "Luco", careers[0].ID),
		models.NewStudent("Sebastian", "Flores", careers[0].ID),
		models.NewStudent("Fabian", "Alvarado", careers[0].ID),

		models.NewStudent("Francisco", "Ibacache", careers[0].ID),
		models.NewStudent("Jaime", "Gatica", careers[0].ID),
		models.NewStudent("Jacob", "Romero", careers[0].ID),

		models.NewStudent("Mario", "Albornoz", careers[0].ID),
		models.NewStudent("Samuel", "Perez", careers[0].ID),
		models.NewStudent("Valentina", "Tarifeño", careers[0].ID),

		models.NewStudent("Cristian", "Reveco", careers[0].ID),
		models.NewStudent("Jordan", "Porras", careers[0].ID),
		models.NewStudent("Felipe", "Campos", careers[0].ID),

		models.NewStudent("Ruben", "Gazitua", careers[0].ID),
		models.NewStudent("Joaquin", "Dinamarca", careers[0].ID),
		models.NewStudent("Esteban", "Lundin", careers[0].ID),

		models.NewStudent("Alberto", "Salinas", careers[0].ID),
		models.NewStudent("Rodrigo", "Tapia", careers[0].ID),
		models.NewStudent("Marcelo", "Araya", careers[0].ID),

		models.NewStudent("Nicolas", "Perez", careers[0].ID),
		models.NewStudent("Sebastian", "Olivares", careers[0].ID),
		models.NewStudent("Marco", "Garrido", careers[0].ID),

		models.NewStudent("Roberto", "Vargas", careers[0].ID),
		models.NewStudent("Felipe", "Fuentes", careers[0].ID),
	}

	students := make([]models.Student, 0)
	for _, st := range sts {
		err := s.Create(&st)
		if err != nil {
			return nil, err
		}
		students = append(students, st)
	}
	return students, nil
}

func generateAdmins(s *services.CRUDService) ([]models.Admin, error) {
	admins := make([]models.Admin, 0)
	admins = append(admins, models.Admin{
		EntryYear: 2015,
		User: models.User{
			Email:     "jacob@utem.cl",
			FirstName: "Jacob",
			LastName:  "Romero",
			Nick:      "jacob.romero",
			RUT:       "19.523.952-5",
			UserType:  "Admin",
		},
	})

	hash, err := admins[0].User.HashPassword("admin123")

	if err != nil {
		return nil, err
	}

	admins[0].User.Password = hash

	for index, admin := range admins {
		admin.InitGSP("access:admin")
		err = s.Create(&admin)
		if err != nil {
			return nil, err
		}
		admins[index] = admin
	}

	return admins, nil
}

func generateTeachers(s *services.CRUDService) ([]models.Teacher, error) {
	teachers := make([]models.Teacher, 0)
	names := []FLName{
		FLName{
			F: "Michael",
			L: "Miranda",
		},
		FLName{
			F: "Oscar",
			L: "Magna",
		},
		FLName{
			F: "Santiago",
			L: "Zapata",
		},
		FLName{
			F: "Victor",
			L: "Escobar",
		},
		FLName{
			F: "Danny",
			L: "Lobos",
		},
		FLName{
			F: "Hector",
			L: "Pincheira",
		},
		FLName{
			F: "Mauro",
			L: "Castillo",
		},
		FLName{
			F: "David",
			L: "Castro",
		},
		FLName{
			F: "Sara",
			L: "Rojas",
		},
		FLName{
			F: "Patricia",
			L: "Mellado",
		},
		FLName{
			F: "Jorge",
			L: "Vergara",
		},
		FLName{
			F: "Cristian",
			L: "Barria",
		},
	}
	for _, pair := range names {
		t := models.NewTeacher(pair.F, pair.L)
		t.User.UserType = "Teacher"
		h, err := t.User.HashPassword(t.User.Nick)
		if err != nil {
			return nil, err
		}
		t.User.Password = h
		err = s.Create(&t)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, t)
	}
	return teachers, nil
}

func generateSubjects(s *services.CRUDService) ([]models.Subject, error) {
	sjs := []models.Subject{
		models.NewSubject("Inteligencia de Negocios", "mdi-chart-bar-stacked"),

		models.NewSubject("Gestión del conocimiento", "mdi-briefcase-account-outline"),
		models.NewSubject("Ingeniería Computacional", "mdi-server"),
		models.NewSubject("Ciencias de la información", "mdi-database"),

		models.NewSubject("Ingeniería de Software", "mdi-code-tags"),
		models.NewSubject("Inteligencia Artificial", "mdi-robot"),
		models.NewSubject("Ciencia de Datos", "mdi-test-tube"),

		models.NewSubject("Ciencias de la computación", "mdi-desktop-classic"),
		models.NewSubject("Seguridad Informática", "mdi-security"),
	}
	subjects := make([]models.Subject, 0)
	for _, sj := range sjs {
		err := s.Create(&sj)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, sj)
	}
	return subjects, nil
}

func generateLinkTypes(s *services.CRUDService) ([]models.LinkType, error) {
	lts := []models.LinkType{
		models.NewLinkType("Normal", "mdi-link"),
		models.NewLinkType("Google Drive", "mdi-folder-google-drive"),
		models.NewLinkType("OneDrive", "mdi-microsoft-onedrive"),
		models.NewLinkType("Repositorio Git", "mdi-git"),
	}
	types := make([]models.LinkType, 0)
	for _, lt := range lts {
		err := s.Create(&lt)
		if err != nil {
			return nil, err
		}
		types = append(types, lt)
	}
	return types, nil
}

func generateProjectTypes(s *services.CRUDService) ([]models.ProjectType, error) {
	ts := []models.ProjectType{
		models.NewProjectType("RESEARCH"),
		models.NewProjectType("PROTOTYPE"),
		models.NewProjectType("APPLICATION"),
		models.NewProjectType("DATA_SCIENCE"),
	}
	types := make([]models.ProjectType, 0)
	for _, t := range ts {
		err := s.Create(&t)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}
	return types, nil
}

func generateProjectStates(s *services.CRUDService) ([]models.ProjectState, error) {
	sts := []models.ProjectState{
		models.NewProjectState("CREATED"),
		models.NewProjectState("STARTED"),
		models.NewProjectState("IN_PROGRESS"),
		models.NewProjectState("PRESENTED"),
		models.NewProjectState("CHECKED"),
		models.NewProjectState("APPROVED"),
		models.NewProjectState("IN_PROGRESS2"),
		models.NewProjectState("PRESENTED2"),
		models.NewProjectState("APPROVED2"),
		models.NewProjectState("REJECTED"),
		models.NewProjectState("FINISHED"),
		models.NewProjectState("CERTIFICATED"),
	}
	states := make([]models.ProjectState, 0)
	for _, st := range sts {
		err := s.Create(&st)
		if err != nil {
			return nil, err
		}
		states = append(states, st)
	}
	return states, nil
}

func generateProjects(s *services.CRUDService, students []models.Student, teachers []models.Teacher, subjects []models.Subject, types []models.ProjectType, states []models.ProjectState) ([]models.Project, error) {
	pjs := []models.Project{
		models.NewProject(
			"Sistema TTS",
			[]models.Student{
				students[0],
			},
			[]models.Teacher{
				teachers[0],
			},
			[]models.Subject{
				subjects[0],
				subjects[1],
			},
			types[0],
		),
		models.NewProject(
			"ChatBot",
			[]models.Student{
				students[1],
			},
			[]models.Teacher{
				teachers[0],
			},
			[]models.Subject{
				subjects[0],
				subjects[1],
			},
			types[2],
		),
		models.NewProject(
			"Sistema TTS",
			[]models.Student{
				students[2],
			},
			[]models.Teacher{
				teachers[0],
			},
			[]models.Subject{
				subjects[0],
				subjects[1],
			},
			types[0],
		),
		models.NewProject(
			"IA",
			[]models.Student{
				students[3],
			},
			[]models.Teacher{
				teachers[0],
			},
			[]models.Subject{
				subjects[0],
				subjects[1],
			},
			types[0],
		),
		models.NewProject(
			"Generación de Redes neuronales",
			[]models.Student{
				students[4],
			},
			[]models.Teacher{
				teachers[0],
			},
			[]models.Subject{
				subjects[2],
				subjects[3],
			},
			types[0],
		),
		models.NewProject(
			"Mesero Virtual",
			[]models.Student{
				students[5],
			},
			[]models.Teacher{
				teachers[0],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Chatbot",
			[]models.Student{
				students[6],
				students[7],
			},
			[]models.Teacher{
				teachers[0],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Chatbot",
			[]models.Student{
				students[8],
			},
			[]models.Teacher{
				teachers[0],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"UTEM SEO",
			[]models.Student{
				students[9],
			},
			[]models.Teacher{
				teachers[1],
			},
			[]models.Subject{
				subjects[1],
				subjects[6],
			},
			types[3],
		),
		models.NewProject(
			"Computación Ubicua",
			[]models.Student{
				students[10],
				students[11],
			},
			[]models.Teacher{
				teachers[2],
			},
			[]models.Subject{
				subjects[1],
			},
			types[0],
		),
		models.NewProject(
			"Neutrosofía y Lógica Difusa",
			[]models.Student{
				students[12],
			},
			[]models.Teacher{
				teachers[2],
			},
			[]models.Subject{
				subjects[1],
			},
			types[0],
		),
		models.NewProject(
			"Inteligencia Ambiental",
			[]models.Student{
				students[13],
			},
			[]models.Teacher{
				teachers[2],
			},
			[]models.Subject{
				subjects[1],
			},
			types[0],
		),
		models.NewProject(
			"Reutiliza UTEM",
			[]models.Student{
				students[14],
			},
			[]models.Teacher{
				teachers[3],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"App Huella de Carbón",
			[]models.Student{
				students[15],
			},
			[]models.Teacher{
				teachers[3],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Mejora Plataforma Sustenta",
			[]models.Student{
				students[16],
			},
			[]models.Teacher{
				teachers[3],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"BIM",
			[]models.Student{
				students[17],
			},
			[]models.Teacher{
				teachers[3],
				teachers[4],
			},
			[]models.Subject{
				subjects[3],
			},
			types[0],
		),
		models.NewProject(
			"Sistema Alerta Temprana",
			[]models.Student{
				students[18],
			},
			[]models.Teacher{
				teachers[3],
			},
			[]models.Subject{
				subjects[1],
			},
			types[2],
		),
		models.NewProject(
			"Analítica de accidentes de tránsito",
			[]models.Student{
				students[19],
			},
			[]models.Teacher{
				teachers[3],
			},
			[]models.Subject{
				subjects[1],
			},
			types[2],
		),
		models.NewProject(
			"Modelo de predicción precio cobre",
			[]models.Student{
				students[20],
			},
			[]models.Teacher{
				teachers[3],
			},
			[]models.Subject{
				subjects[0],
			},
			types[0],
		),
		models.NewProject(
			"API Gateway microservicios UTEM",
			[]models.Student{
				students[21],
			},
			[]models.Teacher{
				teachers[3],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Evaluación de modelos de representación de matrices dispersas",
			[]models.Student{
				students[22],
			},
			[]models.Teacher{
				teachers[5],
			},
			[]models.Subject{
				subjects[7],
			},
			types[2],
		),
		models.NewProject(
			"Paradigmas de programación y procesadores de lenguajes",
			[]models.Student{
				students[23],
			},
			[]models.Teacher{
				teachers[5],
			},
			[]models.Subject{
				subjects[7],
			},
			types[2],
		),
		models.NewProject(
			"Sistema documental de asignatura",
			[]models.Student{
				students[24],
			},
			[]models.Teacher{
				teachers[5],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Planificación de actividades con recursos limitados",
			[]models.Student{
				students[25],
			},
			[]models.Teacher{
				teachers[5],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Modelo de agrupación de problemas en IA",
			[]models.Student{
				students[26],
			},
			[]models.Teacher{
				teachers[5],
			},
			[]models.Subject{
				subjects[4],
			},
			types[0],
		),
		models.NewProject(
			"Registro de Proyectos",
			[]models.Student{
				students[27],
			},
			[]models.Teacher{
				teachers[6], teachers[7],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Sistema Hoja de Vida Estudiantil",
			[]models.Student{
				students[28],
			},
			[]models.Teacher{
				teachers[6],
				teachers[7],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Sistema de Evaluación Por Pares",
			[]models.Student{
				students[29],
			},
			[]models.Teacher{
				teachers[6],
				teachers[7],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Sistema de Administración de Estacionamiento",
			[]models.Student{
				students[30],
			},
			[]models.Teacher{
				teachers[6],
				teachers[7],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Sistema de Control de acceso a instalaciones",
			[]models.Student{
				students[31],
			},
			[]models.Teacher{
				teachers[6],
				teachers[7],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Plataforma colaborativa para docentes",
			[]models.Student{
				students[32],
			},
			[]models.Teacher{
				teachers[6],
				teachers[7],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"BI para empresa AbInBev",
			[]models.Student{
				students[33],
			},
			[]models.Teacher{
				teachers[8],
			},
			[]models.Subject{
				subjects[0],
				subjects[1],
			},
			types[2],
		),
		models.NewProject(
			"Seguridad informática para aplicaciones UTEM",
			[]models.Student{
				students[34],
			},
			[]models.Teacher{
				teachers[8],
			},
			[]models.Subject{
				subjects[8],
			},
			types[2],
		),
		models.NewProject(
			"UTEM Market WEB",
			[]models.Student{
				students[35],
			},
			[]models.Teacher{
				teachers[8],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Report de Añertas",
			[]models.Student{
				students[36],
			},
			[]models.Teacher{
				teachers[8],
			},
			[]models.Subject{
				subjects[0],
			},
			types[2],
		),
		models.NewProject(
			"Plataforma de gestión de aprendizaje",
			[]models.Student{
				students[37],
			},
			[]models.Teacher{
				teachers[9],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Juego Educativo Cultura Chinchorro",
			[]models.Student{
				students[38],
			},
			[]models.Teacher{
				teachers[9],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Gestión Administrativa para Liceo",
			[]models.Student{
				students[39],
			},
			[]models.Teacher{
				teachers[9],
			},
			[]models.Subject{
				subjects[4],
			},
			types[2],
		),
		models.NewProject(
			"Clasificación de Eventos Astronómicos",
			[]models.Student{
				students[40],
			},
			[]models.Teacher{
				teachers[10],
			},
			[]models.Subject{
				subjects[6],
			},
			types[0],
		),
		models.NewProject(
			"Reconocimiento de Hojas de plantas con Deep Learning",
			[]models.Student{
				students[41],
			},
			[]models.Teacher{
				teachers[10],
			},
			[]models.Subject{
				subjects[6],
			},
			types[0],
		),
		models.NewProject(
			"One Shot Learning",
			[]models.Student{
				students[42],
			},
			[]models.Teacher{
				teachers[10],
			},
			[]models.Subject{
				subjects[6],
			},
			types[0],
		),
		models.NewProject(
			"Transfer Learning para reconocimiento de objetos y personas",
			[]models.Student{
				students[43],
			},
			[]models.Teacher{
				teachers[10],
			},
			[]models.Subject{
				subjects[6],
			},
			types[0],
		),
		models.NewProject(
			"Mitigación de vulneración de datos en dispositivos IoT",
			[]models.Student{
				students[44],
			},
			[]models.Teacher{
				teachers[11],
			},
			[]models.Subject{
				subjects[8],
			},
			types[0],
		),
	}
	projects := make([]models.Project, 0)
	for _, p := range pjs {
		p.ProjectStateID = states[2].ID
		err := s.Create(&p)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func generateChannels(s *services.CRUDService) ([]models.Channel, error) {
	chs := []models.Channel{
		models.NewChannel("ZOOM", "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.macupdate.com%2Fimages%2Ficons512%2F52421.png&f=1&nofb=1", true),
		models.NewChannel("Microsoft Teams", "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.bemidjistate.edu%2Foffices%2Fits%2Fwp-content%2Fuploads%2Fsites%2F60%2F2018%2F03%2Ficon-microsoft-teams.png&f=1&nofb=1", true),
		models.NewChannel("Google Hangouts", "https://upload.wikimedia.org/wikipedia/commons/thumb/e/ee/Hangouts_icon.svg/1200px-Hangouts_icon.svg.png", true),
		models.NewChannel("Google Meet", "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.macupdate.com%2Fimages%2Ficons512%2F52421.png&f=1&nofb=1", true),
		models.NewChannel("Generico", "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fimage.flaticon.com%2Ficons%2Fpng%2F512%2F27%2F27130.png&f=1&nofb=1", true),
		models.NewChannel("Presencial", "https://image.flaticon.com/icons/svg/2922/2922328.svg", false),
	}
	channels := make([]models.Channel, 0)
	for _, ch := range chs {
		err := s.Create(&ch)
		if err != nil {
			return nil, err
		}
		channels = append(channels, ch)
	}
	return channels, nil
}

func generateRubric(s *services.CRUDService) ([]models.Rubric, error) {
	rbs := []models.Rubric{
		models.NewRubric(
			"Rúbrica 1",
			"https://drive.google.com/file/d/1uoW8zmfESz0uG6eDvAtxKI3AZmhZsRTk/view?usp=sharing",
		),
		models.NewRubric(
			"Rúbrica 2",
			"https://drive.google.com/file/d/111FX7U9ea6p9W621GiKGw9nDMLL81gpc/view?usp=sharing",
		),
	}
	rubrics := make([]models.Rubric, 0)
	for _, r := range rbs {
		err := s.Create(&r)
		if err != nil {
			return nil, err
		}
		rubrics = append(rubrics, r)
	}
	return rubrics, nil
}

func fillProject(s *services.CRUDService, p *models.Project, rubrics []models.Rubric, channels []models.Channel, ltypes []models.LinkType) error {
	rws := []models.Review{
		models.NewReview("Evaluación 1", rubrics[0].ID, p.ID, "https://drive.google.com/file/d/1KsNk_F7mMMFta6pPXfKFy4uDKpg8kevj/view?usp=sharing", p.Guides[0].ID, "5.4"),
		models.NewReview("Evaluación 2", rubrics[1].ID, p.ID, "https://drive.google.com/file/d/1clFlVTyikx10x_5q3zSANIwgqM3LE7I0/view?usp=sharing", p.Guides[0].ID, "5.5"),
	}
	p.Reviews = rws
	chindex := rand.Intn(len(channels))
	chid := channels[chindex].ID
	mts := []models.Meet{
		models.NewMeet("Reunión #1", time.Now().Add(-1*time.Hour*100), chid, p.ID),
		models.NewMeet("Reunión #2", time.Now().Add(-1*time.Hour*48), chid, p.ID),
		models.NewMeet("Reunión #3", time.Now().Add(time.Hour*72), chid, p.ID),
	}
	p.Meets = mts
	mils := []models.Milestone{
		models.NewMilestone("PreAvance", time.Now().Add(-1*time.Hour*480), p.ID),
		models.NewMilestone("Avance #1", time.Now().Add(-1*time.Hour*240), p.ID),
		models.NewMilestone("Avance #2", time.Now().Add(-1*time.Hour*120), p.ID),
		models.NewMilestone("Entrega Final", time.Now().Add(time.Hour*60), p.ID),
	}
	p.Milestones = mils

	comms := []models.Commit{
		models.NewCommit("Acuerdo #1", time.Now().Add(time.Hour*48), p.ID),
		models.NewCommit("Acuerdo #2", time.Now().Add(time.Hour*24), p.ID),
		models.NewCommit("Acuerdo #3", time.Now().Add(time.Hour*56), p.ID),
		models.NewCommit("Acuerdo #4", time.Now().Add(time.Hour*72), p.ID),
	}
	p.Commits = comms

	pgs := []models.Progress{
		models.NewProgress("Corrección de Cronograma", p.ID),
		models.NewProgress("Corrección ortográfica", p.ID),
	}
	p.Progress = pgs

	links := []models.Link{
		models.NewLink("https://github.com/Moletastic/geopath", ltypes[3].ID, p.ID),
		models.NewLink("https://drive.google.com/drive/u/1/folders/1ZHomNQcMM1C5s3sG5wHhDrh4VXigNRA-", ltypes[1].ID, p.ID),
	}
	p.Links = links
	return s.Update(p, p.ID)
}
