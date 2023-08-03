# Sigen
#### A simple code generator for developers. A tool that allows you to create new codes based on your own templates.

## Usage

### hello world
```shell
make build

./sigen -t repository _pkg=user _Pkg=User
```

In the first step, you need to import your templates into the configuration file. In the configuration file, you can define multiple templates, and when using the tool, you can specify your desired template. Additionally, you can assign values to the variables present in the template.
### Simple
```shell
sigen -t <template> 
```

### Custom config
```shell
sigen -c <path_to_config.yaml> -t <template> 
```

### Custom output path
```shell
sigen -t <template> -o <output_path> 
```

### Assign values to variables
```shell
sigen -t <template> _<var_name>=<value>
````

## Configuration
Let's start with having an agreement on definitions. Directory and link are specific types of files. So, in some sections of the document, the term "file" refers to all three types: files, directories, and links.

In your settings file, you input the structure of your project's files and directories in the form of a standard template. The template is essentially a file and must be one of the three types: file, directory, or link.

The settings of this tool are stored in the YAML format. An example of it is available in the config.yaml file.
### Structure
```yaml
templates:
  <templateName>:
    name: "<file name>"
    type: "<enum:file|directory|link>"
    content: "" // only file|link support this
    sub:
    - name: ""
      type: ""
      .
      .
      .
```

## Template
### Structure

- **name** original name of file without any path, include extension of file.
- **type** an enum to type of file and must be one of the three types: "file", "directory", or "link".
- **content** in "file" type is content of file. in "link" type is destination path of link. (in "directory" type, has not any action)
- **sub** in "directory" type is slice of sub templates, means any sub created into this directory and can be nested continue. (in "file" and "link" type, has not any action)

## Variable
One of the most important features of this tool is the use of variables
in templates. You can use variables in the name or content of the
template file and, each time you invoke the tool, assign values to 
the defined variables.

A variable is defined in the settings file with the format ```${variable_name}```
and during the invocation, it is assigned a value as an argument 
with the format ```_variable_name=value```