package pokeapi

import (
	"example.com/codelytv/cobra-03/cmd/ports"
	"reflect"
	"testing"
)

func TestNewPokeApiRepository(t *testing.T) {
	tests := []struct {
		name string
		want ports.GetPokemonsRepo
	}{
		{name: "Prueba success", want: NewPokeApiRepository()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPokeApiRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPokeApiRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pokeApiRepo_Execute(t *testing.T) {
	type fields struct {
		url string
	}
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantPokemons []ports.Pokemon
		wantErr      bool
	}{
		{name: "Prueba success",
			fields: fields{url: "https://pokeapi.co/api/v2"},
			args: args{limit: 1, offset: 0}, wantPokemons:
				[]ports.Pokemon{{Number:1, Name:"bulbasaur", Height: 7, Weight: 69}},
				wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pokeApiRepo{
				url: tt.fields.url,
			}
			gotPokemons, err := p.Execute(tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPokemons, tt.wantPokemons) {
				t.Errorf("Execute() gotPokemons = %v, want %v", gotPokemons, tt.wantPokemons)
			}
		})
	}
}

func BenchmarkExecute(b *testing.B) {
	type fields struct {
		url string
	}
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
	}{
		{name: "Benchmarking",
			fields: fields{url: "https://pokeapi.co/api/v2"},
			args: args{limit: 100, offset: 0}},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			p := &pokeApiRepo{
				url: tt.fields.url,
			}
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				p.Execute(tt.args.limit, tt.args.offset)
			}
		})
	}
}