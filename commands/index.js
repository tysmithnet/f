import React from 'react'
import { Form, Field } from 'react-final-form'
import { AppContext, Box, Color, Text } from 'ink'
import TextInput from "ink-text-input";
import SelectInput from '../components/SelectInput'
import MultiSelectInput from '../components/MultiSelectInput'
import Error from '../components/Error'
import semver from 'semver'
import fetch from 'node-fetch'
import Spinner from 'ink-spinner'

const npmCache = {}
const checkPackage = name => {
	if (npmCache[name] === undefined) {
		return fetch(`https://api.npms.io/v2/package/${name}`)
			.then(response => response.json())
			.then(json => {
				npmCache[name] = json.code !== 'NOT_FOUND'
				return npmCache[name]
			})
	}
	return npmCache[name]
}

const fields = [
	{
		name: 'name',
		label: 'Project Name',
		validate: value => {
			if (!value) {
				return 'Required'
			}
			const check = checkPackage(value)
			if (check && check.then) {
				return check.then(exists =>
					exists ? 'Package exists already!' : undefined
				)
			}
		},
		format: value =>
			value
				? value
					.toLowerCase()
					.replace(/[^a-z \\-]/g, '')
					.replace(/ /g, '-')
				: '',
		placeholder: 'my-awesome-project',
		Input: TextInput
	},
	{
		name: 'version',
		label: 'Version',
		placeholder: '1.0.0',
		format: value => (value === undefined ? '' : value.replace(/[^0-9.]/g, '')),
		validate: value =>
			!value
				? 'Required'
				: semver.valid(value)
					? undefined
					: 'Invalid semantic version',
		Input: TextInput
	},
	{
		name: 'language',
		label: 'Language',
		Input: SelectInput,
		inputConfig: {
			items: [
				{ label: 'Javascript', value: 'javascript' },
				{ label: 'Typescript', value: 'typescript' }
			]
		}
	},
	{
		name: 'technologies',
		label: 'Technologies',
		Input: MultiSelectInput,
		format: null, // prevents empty value from being ''
		inputConfig: {
			items: [
				{ label: '⚛️ React', value: 'react' },
				{ label: 'Angular', value: 'angular' },
				{ label: 'Redux', value: 'redux' },
				{ label: 'GraphQL', value: 'graphql' },
				{ label: '🏁 React-Final-Form', value: 'react-final-form' },
				{ label: '💅 Styled Components', value: 'styled-components' },
				{ label: '👨‍🎤 Emotion', value: 'emotion' },
				{ label: '🌈‍ Ink', value: 'ink' }
			]
		}
	}
]

const values = "db0,cache1,html,web0,web1,web3,cache0,dbright,tsmith,builduser,dwilson".split(",");


/// CliForm
export default function CliForm() {
	const [searchValue, setSearchValue] = React.useState("");
	const toList = values.filter(x => x.indexOf(searchValue) > -1).map(x => {
		return <Box key={x}>{x}</Box>;
	})
	return (
		<React.Fragment>
			<Box>
				Input: <TextInput value={searchValue} onChange={setSearchValue} />
			</Box>
			<Box>
				{toList}
			</Box>
		</React.Fragment>
	)
}
