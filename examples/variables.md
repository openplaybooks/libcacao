# OpenPlaybooks/libcacao/variables


## Variable Creation

There are three difference ways in Golang that you can create Playbook variables
for CACAO Playbooks.

```
p := playbook.New()
v1 := &objects.Variables{
	ObjectType:  "string",
	Name:        "__test__.Value"
	Description: "FunStuff",
	Value:       "foo",
}
p.AddVariable(*v1)
```

```
p := playbook.New()
v1 := &objects.Variables{
	"string",
	"__test__",
	"Fun Stuff Description",
	"foo",
	false,
	false,
}
p.AddVariable(*v1)
```

```
p := playbook.New()
v1 := objects.NewVariable()
v1.ObjectType = "string"
v1.Name = "__test__.Value"
v1.Description = "Fun Stuff"
v1.Value = "foo"
p.AddVariable(*v1)
```

## License

This is free software, licensed under the Apache License, Version 2.0.
[Read this](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0)) for
a summary.


## Copyright

Copyright 2021 Bret Jordan, All rights reserved.
