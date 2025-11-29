package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Suraj"},
			[]string{"Suraj"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Suraj", "Mumbai"},
			[]string{"Suraj", "Mumbai"},
		},
		{
			"Struct with one string and one int field",
			struct {
				Name string
				Age  int
			}{"Suraj", 22},
			[]string{"Suraj"},
		},
		{
			"Struct with nested fields",
			Person{
				"Suraj",
				Profile{22, "Mumbai"},
			},
			[]string{"Suraj", "Mumbai"},
		},
		{
			"Pointers to nested fields",
			&Person{
				"Suraj",
				Profile{22, "Mumbai"},
			},
			[]string{"Suraj", "Mumbai"},
		},
		{
			"Slices",
			[]Profile{
				{22, "Mumbai"},
				{20, "Pune"},
			},
			[]string{"Mumbai", "Pune"},
		},
		{
			"Arrays",
			[2]Profile{
				{22, "Mumbai"},
				{20, "Pune"},
			},
			[]string{"Mumbai", "Pune"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{22, "Mumbai"}
			aChannel <- Profile{20, "Pune"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Mumbai", "Pune"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
