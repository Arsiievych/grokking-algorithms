package chapter06_breadth_first_search

import (
	"container/list"
	"slices"
)

var UaCityConnection = map[string][]string{
	"Вінниця":          {"Житомир", "Київ", "Черкаси", "Кропивницький", "Одеса", "Хмельницький"},
	"Луцьк":            {"Рівне", "Львів"},
	"Дніпро":           {"Полтава", "Харків", "Донецьк", "Запоріжжя", "Херсон", "Миколаїв", "Кропивницький"},
	"Донецьк":          {"Харків", "Дніпро", "Запоріжжя", "Луганськ"},
	"Житомир":          {"Рівне", "Хмельницький", "Вінниця", "Київ"},
	"Ужгород":          {"Львів", "Івано-Франківськ"},
	"Запоріжжя":        {"Дніпро", "Донецьк", "Херсон"},
	"Івано-Франківськ": {"Львів", "Тернопіль", "Чернівці", "Ужгород"},
	"Київ":             {"Житомир", "Чернігів", "Полтава", "Черкаси", "Вінниця"},
	"Кропивницький":    {"Черкаси", "Полтава", "Дніпро", "Миколаїв", "Одеса", "Вінниця"},
	"Луганськ":         {"Харків", "Донецьк"},
	"Львів":            {"Луцьк", "Рівне", "Тернопіль", "Івано-Франківськ", "Ужгород"},
	"Миколаїв":         {"Одеса", "Кропивницький", "Дніпро", "Херсон"},
	"Одеса":            {"Вінниця", "Кропивницький", "Миколаїв"},
	"Полтава":          {"Суми", "Харків", "Дніпро", "Кропивницький", "Черкаси", "Київ"},
	"Рівне":            {"Луцьк", "Львів", "Тернопіль", "Хмельницький", "Житомир"},
	"Суми":             {"Чернігів", "Полтава", "Харків"},
	"Тернопіль":        {"Львів", "Рівне", "Хмельницький", "Івано-Франківськ"},
	"Харків":           {"Суми", "Полтава", "Дніпро", "Донецьк", "Луганськ"},
	"Херсон":           {"Дніпро", "Миколаїв", "Запоріжжя"},
	"Хмельницький":     {"Рівне", "Житомир", "Вінниця", "Чернівці", "Тернопіль"},
	"Черкаси":          {"Київ", "Вінниця", "Кропивницький", "Полтава"},
	"Чернівці":         {"Івано-Франківськ", "Тернопіль", "Хмельницький"},
	"Чернігів":         {"Київ", "Суми"},
}

func FindShorterWay(from, to string) []string {
	if early, proceed := checkTrivialAndExistence(from, to); !proceed {
		return early
	}

	parents, found := bfsParents(UaCityConnection, from, to)
	if !found {
		return []string{}
	}

	return buildPath(parents, from, to)
}

func checkTrivialAndExistence(from, to string) (early []string, proceed bool) {
	if from == to {
		return []string{from}, false
	}
	if _, ok := UaCityConnection[from]; !ok {
		return []string{}, false
	}
	if _, ok := UaCityConnection[to]; !ok {
		return []string{}, false
	}
	return nil, true
}

func bfsParents(graph map[string][]string, start, target string) (map[string]string, bool) {
	visited := make(map[string]bool)
	parent := make(map[string]string)

	q := list.New()
	q.PushBack(start)
	visited[start] = true

	for q.Len() > 0 {
		node := q.Front()
		current := node.Value.(string)
		q.Remove(node)

		if current == target {
			return parent, true
		}

		for _, neighbor := range graph[current] {
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			parent[neighbor] = current
			q.PushBack(neighbor)
		}

	}

	return parent, false
}

func buildPath(parent map[string]string, from, to string) []string {
	path := []string{}
	for at := to; ; {
		path = append(path, at)
		if at == from {
			break
		}
		p, ok := parent[at]
		if !ok {
			return []string{}
		}
		at = p
	}
	slices.Reverse(path)
	return path
}
